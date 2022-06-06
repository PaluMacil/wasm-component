package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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
		//time from when the connection is accepted to when the request body is fully read
		ReadTimeout: 5 * time.Second,
		//time from the end of the request header read to the end of the response write
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, os.Kill)
	go func() {
		logger.Info("shutting down", zap.Error(srv.ListenAndServe()))
	}()
	logger.Info("now serving", zap.String("port", port))
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
