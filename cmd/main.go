package main

import (
	"context"
	"crypto/tls"
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
	cert, err := tls.LoadX509KeyPair("./cert/localhost-client.pem", "./cert/localhost-client-key.pem")
	if err != nil {
		log.Printf(":: %v", cert)
		return nil, err
	}

	/*
		caCert, err := ioutil.ReadFile("./cert/ExampleCA.pem")
		if err != nil {
			log.Printf("error %v", caCert)
			return nil, err
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
	*/
	//create mongodb connection
	mc, err := mongodb.NewMongoClient()
	if err != nil {
		return nil, err
	}
	//create elasticsearch connection
	es, err := elastic.NewElasticClient(context.Background())
	if err != nil {
		return nil, err
	}
	//create kafka consumer
	kc := pubsub.NewReader()

	srv, err := newServer(serverConfig{
		MongoDB:       mc,
		ElasticSearch: es,
		Kafka:         kc,
		TLS: &tls.Config{
			Certificates: []tls.Certificate{cert},
			//RootCAs:      caCertPool,
		},
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
		//http2Server := http2.Server{}
		//_ = http2.ConfigureServer(srv, &http2Server)
		log.Println("Server Listen And Serve")
		if err := srv.ListenAndServeTLS("./cert/localhost.pem", "./cert/localhost-key.pem"); err != nil {
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
	TLS           *tls.Config
}

func getTLSConfig() *tls.Config {
	//caCert, err := ioutil.ReadFile("/Users/eliasschmoelz/Library/Application Support/mkcert/rootCA.pem")
	//if err != nil {
	//	log.Fatal("Error opening cert file", ", error ", err)
	//}
	//caCertPool := x509.NewCertPool()
	//caCertPool.AppendCertsFromPEM(caCert)
	cert, err := tls.LoadX509KeyPair("./cert/localhost.pem", "./cert/localhost-key.pem")
	if err != nil {
		log.Fatalln(err)
	}
	return &tls.Config{
		//ServerName: host,
		// ClientAuth: tls.NoClientCert,				// Client certificate will not be requested and it is not required
		// ClientAuth: tls.RequestClientCert,			// Client certificate will be requested, but it is not required
		// ClientAuth: tls.RequireAnyClientCert,		// Client certificate is required, but any client certificate is acceptable
		// ClientAuth: tls.VerifyClientCertIfGiven,		// Client certificate will be requested and if present must be in the server's Certificate Pool
		// ClientAuth: tls.RequireAndVerifyClientCert,	// Client certificate will be required and must be present in the server's Certificate Pool
		//ClientAuth: certOpt,
		//ClientCAs:  caCertPool,
		//NextProtos:   []string{"h2"},
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12, // TLS versions below 1.2 are considered insecure - see https://www.rfc-editor.org/rfc/rfc7525.txt for details
	}
}

func newServer(c serverConfig) (*http.Server, error) {
	return &http.Server{
		Addr:         ":8090",
		Handler:      routes(c),
		TLSConfig:    getTLSConfig(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}, nil
}
