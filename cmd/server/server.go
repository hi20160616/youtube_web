package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hi20160616/youtube_web/configs"
	"github.com/hi20160616/youtube_web/internal/server"
	"golang.org/x/sync/errgroup"
)

var (
	address string = configs.Value.Address
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// Web service
	s, err := server.NewServer(address)
	if err != nil {
		log.Println(err)
	}
	g.Go(func() error {
		log.Printf("[%s] Server start on %s", configs.Value.Title, address)
		return s.Start(ctx)
	})
	g.Go(func() error {
		defer log.Printf("[%s] Server stop done.", configs.Value.Title)
		<-ctx.Done() // wait for stop signal
		log.Printf("[%s] Server stop now...", configs.Value.Title)
		return s.Stop(ctx)
	})

	// Update service
	jobs := &server.JobService{}
	g.Go(func() error {
		log.Printf("[%s] Jobs start ...", configs.Value.Title)
		return jobs.Start(ctx)
	})
	g.Go(func() error {
		defer log.Printf("[%s] Jobs stop done.", configs.Value.Title)
		<-ctx.Done() // wait for stop signal
		log.Printf("[%s] Jobs stop now...", configs.Value.Title)
		return jobs.Stop(ctx)
	})

	// Elegant stop
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	g.Go(func() error {
		select {
		case sig := <-sigs:
			fmt.Println()
			log.Printf("[%s] signal caught: %s, ready to quit...", configs.Value.Title, sig.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			log.Printf("[%s] not canceled by context: %s", configs.Value.Title, err)
		} else {
			log.Println(err)
		}
	}
}
