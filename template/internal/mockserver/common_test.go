package mockserver

import (
	"context"
	"github.com/razielsd/rzgrpcmock/server/internal/config"
	"github.com/razielsd/rzgrpcmock/server/internal/logger"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createServer(t *testing.T) *Server {
	cfg := &config.Config{}
	api := NewApiServer(cfg, logger.TestLogger(t))
	return api
}

func createGetReqAndWriter() (*httptest.ResponseRecorder, *http.Request) {
	r, _ := http.NewRequestWithContext(
		context.Background(), http.MethodGet, "/", nil,
	)
	w := httptest.NewRecorder()
	return w, r
}