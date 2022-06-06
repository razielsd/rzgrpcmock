package reqmatcher

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.uber.org/zap"
)

const (
	DefaultMatcher             = "default"
	MetaKey        MetaKeyName = "meta"
)

type MetaKeyName string

type RequestMeta struct {
	Method string
}

type MatchRule struct {
	Request    string `json:"request"`
	Response   string `json:"response"`
	MethodName string `json:"method_name"`
}

type Matcher struct {
	ruleList []*MatchRule
	mu       *sync.RWMutex
	log      *zap.Logger
}

var matcherList map[string]*Matcher
var ErrRuleNotFound = status.Error(codes.FailedPrecondition, "rule not found")
var ErrFailedUnmarshal = status.Error(codes.Internal, "failed unmarshal response")

func init() {
	matcherList = make(map[string]*Matcher)
}

func NewMatcher(name string, log *zap.Logger) *Matcher {
	m, ok := matcherList[name]
	if ok {
		return m
	}
	matcherList[name] = &Matcher{
		ruleList: make([]*MatchRule, 0),
		mu:       &sync.RWMutex{},
		log:      log,
	}
	return matcherList[name]
}

func GetMatcher(name string) *Matcher {
	m := matcherList[name]
	return m
}

func (m *Matcher) Match(ctx context.Context, req, resp interface{}) error {
	meta, ok := ctx.Value(MetaKey).(RequestMeta)
	if !ok {
		m.log.Error(
			"failed extract request meta",
			zap.String("method", meta.Method),
		)
		return errors.New("failed extract request meta")
	}
	m.log.Info(
		"Match request",
		zap.String("method", meta.Method),
	)
	var eqRule *MatchRule
	eqWeight := 0
	for _, rule := range m.ruleList {
		weight, isEqual := m.isEqual(rule, meta, req)
		if isEqual && weight > eqWeight {
			eqWeight = weight
			eqRule = rule
		}
	}
	if eqRule == nil {
		return ErrRuleNotFound
	}
	err := json.Unmarshal([]byte(eqRule.Response), &resp)
	if err != nil {
		return ErrFailedUnmarshal
	}
	m.log.Info(
		"Match success",
		zap.String("method", meta.Method),
	)
	return nil
}

func (m *Matcher) Append(rule *MatchRule) {
	m.log.Info(
		"add rule", zap.String("method", rule.MethodName),
	)
	m.mu.Lock()
	defer m.mu.Unlock()
	rule.MethodName = strings.ToLower(strings.TrimSpace(rule.MethodName))
	m.ruleList = append(m.ruleList, rule)
}

func (m *Matcher) Reset() {
	m.log.Info("clear all rules")
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ruleList = make([]*MatchRule, 0)
}

func (m *Matcher) isEqual(rule *MatchRule, meta RequestMeta, req interface{}) (int, bool) {
	if rule.MethodName != meta.Method {
		return 0, false
	}
	js, err := json.Marshal(req)
	if err != nil {
		return 0, false
	}
	expected := make(map[string]interface{})
	in := make(map[string]interface{})

	err = json.Unmarshal([]byte(rule.Request), &expected)
	if err != nil {
		return 0, false
	}

	err = json.Unmarshal(js, &in)
	if err != nil {
		return 0, false
	}
	weight := 1
	for k, v := range expected {
		actual, ok := in[k]
		if !ok {
			return 0, false
		}
		if !reflect.DeepEqual(v, actual) {
			return 0, false
		}
		weight++
	}
	return weight, true
}
