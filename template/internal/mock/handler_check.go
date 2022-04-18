package mock

import "net/http"

func (s *Server) handlerTest(w http.ResponseWriter, r *http.Request) {
	s.sendResult(w, "OK")
}
