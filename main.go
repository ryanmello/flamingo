package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ryanmello/product-microservice/handlers"
)

func main(){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 *time.Second,
		ReadTimeout: 1 *time.Second,
		WriteTimeout: 1 *time.Second,
	}

	go func(){
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	
	
	tc := context.WithDeadline(context.Background(), 30 *time.Second)
	s.Shutdown(tc)
}