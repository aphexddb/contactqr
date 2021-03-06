package contactqr

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
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
	log.Printf("Config -> path: %s, index: %s, port: %s\n", staticPath, indexFile, port)

	r := mux.NewRouter()

	// api routes
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Path("/vcard/create").HandlerFunc(CreateVCardHandler).Methods(http.MethodPost)
	api.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	// API Middleware
	api.Use(LoggingMiddleware)

	// health check
	r.HandleFunc("/health", HealthCheckHandler).Methods(http.MethodGet)

	// catch-all: Serve all static HTML files
	r.PathPrefix("/").HandlerFunc(StaticHTMLHandler(staticPath, indexFile)).Methods(http.MethodGet)

	// handle CORS for local dev
	localDev := os.Getenv("LOCALDEV")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"contactqr.me"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	if len(localDev) > 0 {
		log.Println("Running in local dev mode, allowing all CORS requests")
		originsOk = handlers.AllowedOrigins([]string{"*"})
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: httpWriteTimeout,
		ReadTimeout:  httpReadTimeout,
		IdleTimeout:  httpIdleTimeout,
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(r),
	}

	return &contactqrServer{
		port,
		srv,
	}
}

// Start starts the http server
func (s *contactqrServer) Start() error {
	log.Println("HTTP server starting on port", s.port)
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
