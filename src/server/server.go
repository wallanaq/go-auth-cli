package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

var (
	httpServer *http.Server
	done       chan bool = make(chan bool)
)

func Start(addr string, handler http.Handler) {
	httpServer = &http.Server{Addr: addr, Handler: handler}

	go func() {
		log.Printf("Starting server at %s\n", addr)

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
}

func Shutdown() {
	if httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatalf("Error when shutting down the server: %v", err)
		}

		log.Println("Server gracefully stopped")
	}
}

func Done() chan bool {
	return done
}

func NotifyDone() {
	done <- true
}
