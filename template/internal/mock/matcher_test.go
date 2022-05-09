package mock

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetMatcher(t *testing.T) {
}

func TestMatcher_Append(t *testing.T) {
	matcher := NewMatcher(DefaultMatcher, testLogger(t))
	rule := &MatchRule{
		ServiceName: "s1",
		MethodName:  "m1",
		Request: `{
			  "id": "id-1",
			  "name": "username"
			}`,
		Response: `{"id": "id-1"}`,
	}
	matcher.Append(rule)
	require.Len(t, matcher.ruleList, 1)
	require.Equal(t, rule, matcher.ruleList[0])
}

func TestMatcher_Match(t *testing.T) {
	t.Run("success match rule with 2 matched rule", func(t *testing.T) {
		matcher := NewMatcher(DefaultMatcher, testLogger(t))
		serviceName := "s2"
		methodName := "m2"
		rule := &MatchRule{
			ServiceName: serviceName,
			MethodName:  methodName,
			Request: `{
			  "id": "id-1",
			  "name": "username"
			}`,
			Response: `{"id": "id-1"}`,
		}
		rule2 := &MatchRule{
			ServiceName: serviceName,
			MethodName:  methodName,
			Request: `{
			  "id": "id-1",
			  "name": "username",
              "some_key": "123"
			}`,
			Response: `{"id": "id-2"}`,
		}
		matcher.Append(rule)
		matcher.Append(rule2)
		req := map[string]string{
			"id": "id-1",
			"name": "username",
			"some_key": "123",
		}
		resp := make(map[string]string)
		err := matcher.Match(serviceName, methodName, &req, &resp)
		require.NoError(t, err)

		expected := map[string]string{
			"id": "id-2",
		}
		require.Equal(t, expected, resp)
	})
}

func TestMatcher_isEqual(t *testing.T) {
	t.Run("service name is invalid", func(t *testing.T) {
		matcher := NewMatcher(DefaultMatcher, testLogger(t))
		serviceName := "s2"
		methodName := "m2"
		rule := &MatchRule{
			ServiceName: "s1",
			MethodName:  "m2",
		}
		_, f := matcher.isEqual(rule, serviceName, methodName, nil)
		require.False(t, f)
	})

	t.Run("method name is invalid", func(t *testing.T) {
		matcher := NewMatcher(DefaultMatcher, testLogger(t))
		serviceName := "s1"
		methodName := "m2"
		rule := &MatchRule{
			ServiceName: "s1",
			MethodName:  "m1",
		}
		_, f := matcher.isEqual(rule, serviceName, methodName, nil)
		require.False(t, f)
	})

	t.Run("success weight count", func(t *testing.T) {
		matcher := NewMatcher(DefaultMatcher, testLogger(t))
		serviceName := "s1"
		methodName := "m1"
		rule := &MatchRule{
			ServiceName: "s1",
			MethodName:  "m1",
			Request: `{
			  "id": "id-1",
			  "name": "username"
			}`,
		}
		req := map[string]string{
			"id": "id-1",
			"name": "username",
			"some_key": "123",
		}
		weight, f := matcher.isEqual(rule, serviceName, methodName, req)
		require.True(t, f)
		require.Equal(t, 2, weight)
	})
}

func TestNewMatcher(t *testing.T) {
	logger := NewMatcher("test", testLogger(t))
	require.NotNil(t, logger)
}
