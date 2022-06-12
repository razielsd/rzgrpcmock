package mockserver

import (
	"net/http"

	"github.com/razielsd/rzgrpcmock/template/internal/reqmatcher"
)

func (s *Server) handlerMockAdd(w http.ResponseWriter, r *http.Request) {
	form, err := s.getForm(w, r, []string{"method", "request", "response", "ref"})
	if err != nil {
		s.sendError(w, ErrCodeBadRequest, "unable parse request", err)
		return
	}
	matchRule := &reqmatcher.MatchRule{
		Request:    form["request"],
		Response:   form["response"],
		MethodName: form["method"],
	}
	reqmatcher.GetMatcher(reqmatcher.DefaultMatcher).Append(matchRule)
	s.sendResult(w, NewSuccessOK())
}
