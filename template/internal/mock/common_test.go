package mock

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func testLogger(t *testing.T) *zap.Logger {
	log, err := zap.NewDevelopment()
	require.NoError(t, err, "failed init test logger")
	return log
}
