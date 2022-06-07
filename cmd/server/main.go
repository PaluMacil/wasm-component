package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("creating logger: %s", err)
	}
	defer logger.Sync()

	mux := http.NewServeMux()

	port := "5798"
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
		// time from when the connection is accepted to when the request body is fully read
		ReadTimeout: 5 * time.Second,
		// time from the end of the request header read to the end of the response write
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	stop := make(chan os.Signal, 1)
	// SIGHUP  is sent when a tty disconnects, but but convention it is also commonly used to request that a daemon
	//         reload its configuration
	// SIGINT  is for ctrl-c in the terminal (and cannot be sent programmatically on Windows)
	// SIGTERM is sent by Kubernetes, Docker, and others. Kubernetes gives 30 seconds to exit by default. The closest
	//         Windows equivalent is WM_CLOSE which is sent to a windowed application when the x button is clicked, etc.
	// SIGKILL cannot be delayed, handled, or ignored. This is how you kill a process immediately.
	// 	       On Windows, which doesn't have signals, the equivalent is TerminateProcess (processthreadsapi.h)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		logger.Info("shutting down", zap.Error(srv.ListenAndServe()))
	}()
	logger.Info("now serving", zap.String("port", port))
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Info("shutdown failed", zap.Error(err))
	}
}
