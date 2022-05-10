package mockserver

import (
	"github.com/razielsd/rzgrpcmock/server/internal/reqmatcher"
	"net/http"
)

func (s *Server) handlerMockAdd(w http.ResponseWriter, r *http.Request) {
	form, err := s.getForm(w, r, []string{"service_name", "method", "request", "response", "ref"})
	if err != nil {
		s.sendError(w, ErrCodeBadRequest, "uanble parse request", err)
		return
	}
	matchRule := &reqmatcher.MatchRule{
		Request:     form["request"],
		Response:    form["response"],
		ServiceName: form["service_name"],
		MethodName:  form["method"],
	}
	reqmatcher.GetMatcher(reqmatcher.DefaultMatcher).Append(matchRule)
	s.sendResult(w, NewSuccessOK())
}
