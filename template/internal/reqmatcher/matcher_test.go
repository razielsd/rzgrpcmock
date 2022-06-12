package reqmatcher

import (
	"context"
	"testing"

	"github.com/razielsd/rzgrpcmock/template/internal/logger"
	"github.com/stretchr/testify/require"
)

func TestGetMatcher(t *testing.T) {
}

func TestMatcher_Append(t *testing.T) {
	matcher := NewMatcher(DefaultMatcher, logger.TestLogger(t))
	rule := &MatchRule{
		MethodName: "m1",
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
		matcher := NewMatcher(DefaultMatcher, logger.TestLogger(t))
		methodName := "m2"
		rule := &MatchRule{
			MethodName: methodName,
			Request: `{
			  "id": "id-1",
			  "name": "username"
			}`,
			Response: `{"id": "id-1"}`,
		}
		rule2 := &MatchRule{
			MethodName: methodName,
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
			"id":       "id-1",
			"name":     "username",
			"some_key": "123",
		}
		resp := make(map[string]string)
		meta := RequestMeta{Method: methodName}
		ctx := context.WithValue(context.Background(), MetaKey, meta)
		err := matcher.Match(ctx, &req, &resp)
		require.NoError(t, err)

		expected := map[string]string{
			"id": "id-2",
		}
		require.Equal(t, expected, resp)
	})
}

func TestMatcher_isEqual(t *testing.T) {
	t.Run("method name is invalid", func(t *testing.T) {
		matcher := NewMatcher(DefaultMatcher, logger.TestLogger(t))
		methodName := "m2"
		rule := &MatchRule{
			MethodName: "m1",
		}
		meta := RequestMeta{
			Method: methodName,
		}
		_, f := matcher.isEqual(rule, meta, nil)
		require.False(t, f)
	})

	t.Run("success weight count", func(t *testing.T) {
		matcher := NewMatcher(DefaultMatcher, logger.TestLogger(t))
		methodName := "m1"
		rule := &MatchRule{
			MethodName: "m1",
			Request: `{
			  "id": "id-1",
			  "name": "username"
			}`,
		}
		req := map[string]string{
			"id":       "id-1",
			"name":     "username",
			"some_key": "123",
		}
		meta := RequestMeta{
			Method: methodName,
		}
		weight, f := matcher.isEqual(rule, meta, req)
		require.True(t, f)
		require.Equal(t, 3, weight)
	})
}

func TestNewMatcher(t *testing.T) {
	logger := NewMatcher("test", logger.TestLogger(t))
	require.NotNil(t, logger)
}
