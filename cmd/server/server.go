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
	log.Default().SetPrefix(configs.Value.Title + ": ")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// Web service
	s, err := server.NewServer(address)
	if err != nil {
		log.Println(err)
	}
	g.Go(func() error {
		log.Printf("Server start on %s", address)
		return s.Start(ctx)
	})
	g.Go(func() error {
		defer log.Printf("Server stop done.")
		<-ctx.Done() // wait for stop signal
		log.Printf("Server stop now...")
		return s.Stop(ctx)
	})

	// Update service
	jobs := &server.JobService{}
	g.Go(func() error {
		log.Printf("Jobs start ...")
		return jobs.Start(ctx)
	})
	g.Go(func() error {
		defer log.Printf("Jobs stop done.")
		<-ctx.Done() // wait for stop signal
		log.Printf("Jobs stop now...")
		return jobs.Stop(ctx)
	})

	// Elegant stop
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	g.Go(func() error {
		select {
		case sig := <-sigs:
			fmt.Println()
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			log.Printf("not canceled by context: %s", err)
		} else {
			log.Println(err)
		}
	}
}
