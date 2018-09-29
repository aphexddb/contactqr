package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aphexddb/contactqr"
	"github.com/gorilla/mux"
)

// Reasonably snappy HTTP timeouts
var (
	httpReadTimeout  = 5 * time.Second
	httpWriteTimeout = 10 * time.Second
	httpIdleTimeout  = 60 * time.Second
)

// grace period for http server to shutdown before exiting
var shutdownTimeout = 5 * time.Second

func main() {
	var staticPath string
	var indexPath string
	var port string

	// service flags
	flag.StringVar(&staticPath, "path", "./dist/static", "The directory to serve static files from.")
	flag.StringVar(&indexPath, "index", "./dist/index.html", "The index file to to serve")
	flag.StringVar(&port, "port", "8080", "HTTP `port`")
	flag.Parse()

	r := mux.NewRouter()

	// api routes
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Path("/vcard/create").HandlerFunc(contactqr.NewVCardHandler).Methods(http.MethodPost)
	api.NotFoundHandler = http.HandlerFunc(contactqr.NotFoundHandler)

	// Serve static assets directly
	s := http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath)))
	r.PathPrefix("/static/").Handler(s).Methods(http.MethodGet)

	// health check
	r.HandleFunc("/health", contactqr.HealthCheckHandler).Methods(http.MethodGet)

	// Catch-all: Serve our JavaScript application's entry-point (index.html)
	r.PathPrefix("/").HandlerFunc(contactqr.IndexHandler(indexPath)).Methods(http.MethodGet)

	// log all requests
	r.Use(contactqr.LoggingMiddleware)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: httpWriteTimeout,
		ReadTimeout:  httpReadTimeout,
		IdleTimeout:  httpIdleTimeout,
		Handler:      r,
	}

	// Run server in a goroutine so that it doesn't block
	go func() {
		log.Println("HTTP starting on port", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// listen for signals and block until recevied
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-stopChan

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shutting down")
	os.Exit(0)
}
