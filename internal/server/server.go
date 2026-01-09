package server

import (
	"log"
	"net/http"
	"time"
)

type handler interface {
	HandleRoot(w http.ResponseWriter, r *http.Request)
	HandleUpload(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	server *http.Server
	logger *log.Logger
}

func New(
	logger *log.Logger,
	addr string,
	port string,
	writeTimeout int, // in seconds
	readTimeout int, // in seconds
	idleTimeout int, // in seconds
	handler handler,
) *Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HandleRoot)
	mux.HandleFunc("/upload", handler.HandleUpload)

	server := &http.Server{
		Handler:      mux,
		Addr:         addr + ":" + port,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		IdleTimeout:  time.Duration(idleTimeout) * time.Second,
		ErrorLog:     logger,
	}

	return &Server{
		server: server,
		logger: logger,
	}
}

func (s *Server) Start() error {
	s.logger.Printf("Starting server on %s\n", s.server.Addr)
	return s.server.ListenAndServe()
}
