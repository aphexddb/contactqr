package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aphexddb/contactqr"
)

func main() {
	var staticPath string
	var indexFile string
	var port string

	// service flags
	flag.StringVar(&staticPath, "path", "./dist/static", "The directory to serve static files from")
	flag.StringVar(&indexFile, "index", "index.html", "The http index file to serve")
	flag.StringVar(&port, "port", "8080", "HTTP `port`")
	flag.Parse()

	// create the server
	server := contactqr.NewServer(staticPath, indexFile, port)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// listen for signals and block until recevied
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-stop

	server.Stop()
	os.Exit(0)
}
