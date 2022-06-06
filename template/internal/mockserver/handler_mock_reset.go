package mockserver

import (
	"net/http"

	"github.com/razielsd/rzgrpcmock/server/internal/reqmatcher"
)

func (s *Server) handlerMockReset(w http.ResponseWriter, _ *http.Request) {
	reqmatcher.GetMatcher(reqmatcher.DefaultMatcher).Reset()
	s.sendResult(w, NewSuccessOK())
}
