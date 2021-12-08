package main

import (
	"context"
	"github.com/eliassebastian/gor6-api/internal/elastic"
	"github.com/eliassebastian/gor6-api/internal/mongodb"
	"github.com/eliassebastian/gor6-api/internal/pubsub"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	errC, err := run()
	if err != nil {
		log.Fatalf("Error Running - %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Channel Error - %s", err)
	}
}

func run() (<-chan error, error) {
	//create mongodb connection
	mc, err := mongodb.NewMongoClient()
	if err != nil {
		return nil, err
	}
	//create elasticsearch connection
	es, err := elastic.NewElasticClient()
	if err != nil {
		return nil, err
	}
	//create kafka consumer
	kc := pubsub.NewReader()

	srv, err := newServer(serverConfig{
		MongoDB:       mc,
		ElasticSearch: es,
		Kafka:         kc,
	})

	errC := make(chan error, 1)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		log.Println("Shutdown signal received")
		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			kc.Close()
			mc.Close()

			stop()
			cancel()
			close(errC)
		}()

		srv.SetKeepAlivesEnabled(false)

		if err := srv.Shutdown(ctxTimeout); err != nil {
			errC <- err
		}

		log.Println("Shutdown Complete")
	}()

	go kc.Run(ctx)

	go func() {
		log.Println("Server Listen And Serve")
		if err := srv.ListenAndServe(); err != nil {
			errC <- err
		}
	}()

	return errC, nil
}

type serverConfig struct {
	//Address       string
	MongoDB       *mongodb.MongoClient
	ElasticSearch *elastic.ESClient
	Kafka         *pubsub.Consumer
}

func newServer(c serverConfig) (*http.Server, error) {
	return &http.Server{
		Addr:         ":8090",
		Handler:      routes(c),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}, nil
}
