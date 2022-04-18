package mock

import (
	"encoding/json"
	"reflect"
	"strings"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	for _, rule := range m.ruleList {
		if m.isEqual(rule, serviceName, methodName, req) {
			err := json.Unmarshal([]byte(rule.Response), resp)
			if err != nil {
				return status.Error(codes.Internal, "failed unmarshal response")
			}
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

func (m *Matcher) isEqual(rule *MatchRule, serviceName, methodName string, req interface{}) bool {
	if rule.ServiceName != serviceName || rule.MethodName != methodName {
		return false
	}
	js, err := json.Marshal(req)
	if err != nil {
		return false
	}
	expected := make(map[string]interface{})
	in := make(map[string]interface{})

	err = json.Unmarshal([]byte(rule.Request), &expected)
	if err != nil {
		return false
	}

	err = json.Unmarshal(js, &in)
	if err != nil {
		return false
	}

	for k, v := range expected {
		actual, ok := in[k]
		if !ok {
			return false
		}
		if !reflect.DeepEqual(v, actual) {
			return false
		}
	}
	return true
}
