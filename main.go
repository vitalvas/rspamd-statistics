package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	router := bunrouter.New(
		bunrouter.Use(
			reqlog.NewMiddleware(
				reqlog.WithVerbose(true),
			),
		),
	)

	router.GET("/healthz", func(w http.ResponseWriter, req bunrouter.Request) error {
		_, err := w.Write([]byte("ok"))
		return err
	})

	router.Use(contentTypeMiddleware).POST("/request", requestHandler)

	httpServer := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	httpServer.SetKeepAlivesEnabled(true)

	certFile := flag.String("certFile", "", "Certificate pem file")
	keyFile := flag.String("keyFile", "", "Certificate key file")

	flag.StringVar(&httpServer.Addr, "httpAddr", ":8000", "Listen http address")

	flag.Parse()

	go func() {
		var err error

		if certFile != nil && len(*certFile) > 0 && keyFile != nil && len(*keyFile) > 0 {
			err = httpServer.ListenAndServeTLS(*certFile, *keyFile)
		} else {
			err = httpServer.ListenAndServe()
		}

		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	notifyCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-notifyCtx.Done()

	log.Println("Shutdown signal received")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), httpServer.IdleTimeout)
	defer cancel()

	if err := httpServer.Shutdown(timeoutCtx); err != nil {
		log.Fatal(err)
	}
}
