package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jesperkha/w4d/config"
	"github.com/jesperkha/w4d/server"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	s := server.New(config)
	go s.ListenAndServe(ctx, &wg)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

	cancel()
	wg.Wait()

	log.Println("Shutdown")
}
