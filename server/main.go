package main

import (
	"context"
	"ditto/core"
	"ditto/model/config"
	"log"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// Create and initialize core components
	core.Init()

	// New a LineBot
	// line.NewLineBot()

	// Create a HTTP server
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Config.HttpPort),
		Handler: core.Ditto.Engine,
	}

	// Run gRPC server
	// go rpc.Run(config.Config.TcpPort)

	// Run HTTP server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}
}
