package mockserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/razielsd/rzgrpcmock/template/internal/reqmatcher"

	"github.com/razielsd/rzgrpcmock/template/internal/config"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

const ErrCodeBadRequest = 400

type Server struct {
	Addr string
	log  *zap.Logger
}

func NewApiServer(cfg *config.Config, log *zap.Logger) *Server {
	return &Server{
		Addr: cfg.APIAddr,
		log:  log,
	}
}

func (s *Server) Run(ctx context.Context) {
	reqmatcher.NewMatcher(reqmatcher.DefaultMatcher, s.log)
	r := mux.NewRouter()
	s.addRoute(r)
	srv := s.createServer(r)
	s.startServer(srv)
	<-ctx.Done()
	s.stopServer(srv)
	s.log.Info("server stopped")
}

func (s *Server) addRoute(r *mux.Router) {
	r.HandleFunc("/api/mock/add", s.handlerMockAdd).Methods("POST")
	r.HandleFunc("/api/mock/reset", s.handlerMockReset).Methods("POST")
	r.HandleFunc("/api/form", s.handlerForm).Methods("GET")
	r.HandleFunc("/health/liveness", s.handlerHealthProbe).Methods("GET")
	r.HandleFunc("/health/readiness", s.handlerHealthProbe).Methods("GET")
	r.Path("/metrics").Handler(promhttp.Handler())
}

func (s *Server) createServer(r *mux.Router) *http.Server {
	return &http.Server{
		Addr:         s.Addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      newHTTPLog(r, s.log),
	}
}

func (s *Server) startServer(srv *http.Server) {
	go func() {
		s.log.Info("start api server", zap.String("host", s.Addr))
		if err := srv.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				s.log.Info("server closed")
			} else {
				s.log.Error("error in server", zap.Error(err))
			}
		}
	}()
}

func (s *Server) stopServer(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		s.log.Error("error during shutdown", zap.Error(err))
	}
}

func (s *Server) sendResult(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(SuccessResponse{Result: data})
	if err != nil {
		s.log.Error("Unable encode response", zap.Error(err))
	}
}

func (s *Server) getForm(w http.ResponseWriter, r *http.Request, params []string) (map[string]string, error) {
	if err := r.ParseForm(); err != nil {
		return nil, fmt.Errorf("cannot parse post params %w", err)
	}
	result := make(map[string]string)
	for _, key := range params {
		v := r.Form.Get(key)
		result[key] = v
	}
	return result, nil
}

func (s *Server) sendError(w http.ResponseWriter, code int, message string, err error) {
	w.WriteHeader(http.StatusBadRequest)
	encErr := json.NewEncoder(w).Encode(
		ErrorResponse{
			ErrMessage: fmt.Sprintf("%s: %s", message, err),
			Code:       code,
		})
	if encErr != nil {
		s.log.Error("Error encode response", zap.Error(err))
	}
}
