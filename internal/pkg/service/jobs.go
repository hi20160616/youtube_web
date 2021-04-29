package service

import (
	"context"
	"log"

	"github.com/hi20160616/youtube_web/internal/jobs"
)

type JobService struct {
}

func (j *JobService) Start(ctx context.Context) error {
	log.Println("Jobs start ...")
	if err := jobs.UpdateByHoursStart(ctx); err != nil {
		return err
	}
	return ctx.Err()
}

func (j *JobService) Stop(ctx context.Context) error {
	log.Println("Jobs stop ...")
	if err := jobs.UpdateByHoursStop(); err != nil {
		return err
	}
	return ctx.Err()
}
