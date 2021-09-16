package cron

import (
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

func InitCron(f interface{}) {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every("1m").Do(f)
	if err != nil {
		log.Fatalf("error initializing cron %v", err)
	}
	s.StartAsync()
}
