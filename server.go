package contactqr

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Reasonable HTTP timeouts
var (
	httpReadTimeout  = 5 * time.Second
	httpWriteTimeout = 10 * time.Second
	httpIdleTimeout  = 60 * time.Second
)

// grace period for http server to shutdown before exiting
var shutdownTimeout = 5 * time.Second

// Server defines the ContactQR server
type Server interface {
	Start() error
	Stop()
}

// contactqrServer implements the Server interface
type contactqrServer struct {
	port       string
	httpServer *http.Server
}

// NewServer creates a new server
func NewServer(staticPath, indexFile, port string) Server {
	log.Printf("Server config -> path: %s, index: %s, port: %s\n", staticPath, indexFile, port)

	r := mux.NewRouter()

	// api routes
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Path("/vcard/create").HandlerFunc(NewVCardHandler).Methods(http.MethodPost)
	api.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	// log all API requests
	api.Use(LoggingMiddleware)

	// health check
	r.HandleFunc("/health", HealthCheckHandler).Methods(http.MethodGet)

	// Catch-all: Serve all static HTML files
	r.PathPrefix("/").HandlerFunc(StaticHTMLHandler(staticPath)).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: httpWriteTimeout,
		ReadTimeout:  httpReadTimeout,
		IdleTimeout:  httpIdleTimeout,
		Handler:      r,
	}

	return &contactqrServer{
		port,
		srv,
	}
}

// Start starts the http server
func (s *contactqrServer) Start() error {
	log.Println("HTTP server listening on port", s.port)
	return s.httpServer.ListenAndServe()
}

// Stop stops the http server
func (s *contactqrServer) Stop() {
	// use a timeout deadline to try and gracefully shutdown
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	s.httpServer.Shutdown(ctx)
	log.Println("HTTP server shutting down")
}
