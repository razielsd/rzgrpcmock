package reqmatcher

import (
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
	"strings"
	"sync"

	"go.uber.org/zap"
)

type MatchRule struct {
	Request     string `json:"request"`
	Response    string `json:"response"`
	ServiceName string `json:"service_name"`
	MethodName  string `json:"method_name"`
	Id          string
}

const DefaultMatcher = "default"

type Matcher struct {
	ruleList []*MatchRule
	mu       *sync.RWMutex
	log      *zap.Logger
}

var matcherList map[string]*Matcher

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
	m, _ := matcherList[name]
	return m
}

func (m *Matcher) Match(serviceName, methodName string, req, resp interface{}) error {
	m.log.Info(
		"Match",
		zap.String("serviceName", serviceName),
		zap.String("methodName", methodName),
	)
	serviceName = strings.ToLower(strings.TrimSpace(serviceName))
	methodName = strings.ToLower(strings.TrimSpace(methodName))
	var eqRule *MatchRule
	eqWeight := 0
	for _, rule := range m.ruleList {
		weight, isEqual := m.isEqual(rule, serviceName, methodName, req)
		if  isEqual && weight > eqWeight{
			eqWeight = weight
			eqRule = rule
		}
	}
	if eqRule != nil {
		err := json.Unmarshal([]byte(eqRule.Response), &resp)
		if err != nil {
			return status.Error(codes.Internal, "failed unmarshal response")
		}
	}
	return nil
}

func (m *Matcher) Append(rule *MatchRule) {
	m.log.Info(
		"add rule", zap.String("serviceName", rule.ServiceName), zap.String("method", rule.MethodName),
	)
	m.mu.Lock()
	defer m.mu.Unlock()
	rule.ServiceName = strings.ToLower(strings.TrimSpace(rule.ServiceName))
	rule.MethodName = strings.ToLower(strings.TrimSpace(rule.MethodName))
	m.ruleList = append(m.ruleList, rule)
}

func (m *Matcher) isEqual(rule *MatchRule, serviceName, methodName string, req interface{}) (int, bool) {
	if rule.ServiceName != serviceName || rule.MethodName != methodName {
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
	weight := 0
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
