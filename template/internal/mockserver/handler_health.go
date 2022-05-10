package mockserver

import (
	"net/http"
)

func (s *Server) handlerHealthProbe(w http.ResponseWriter, r *http.Request) {
	s.sendResult(w, NewSuccessOK())
}
