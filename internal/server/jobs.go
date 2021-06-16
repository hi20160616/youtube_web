package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hi20160616/youtube_web/configs"
	"github.com/hi20160616/youtube_web/internal/jobs"
	"github.com/robfig/cron/v3"
)

type JobService struct {
	c *cron.Cron
}

var done = make(chan struct{}, 1)

// Start use cron to invoke func
func (j *JobService) Start(ctx context.Context) error {
	do := func() {
		if err := jobs.UpdateActions(ctx); err != nil {
			log.SetPrefix(configs.Value.Title)
			log.Printf("UpdateActions: %v", err)
		}
	}

	j.c = cron.New(
		cron.WithLogger(
			cron.VerbosePrintfLogger(log.New(os.Stdout,
				"cron: ", log.LstdFlags))))
	j.c.AddFunc(configs.Value.Cron, do)

	do() // do just started

	j.c.Start()
	return ctx.Err()
}

func (j *JobService) Stop(ctx context.Context) error {
	ctx = j.c.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(30 * time.Second):
		return fmt.Errorf("context was not done immediately")
	}
}
