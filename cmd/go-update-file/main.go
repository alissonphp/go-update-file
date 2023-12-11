package main

import (
	"fmt"
	"github.com/alissonphp/go-update-file/infrastructure/router"
	"github.com/alissonphp/go-update-file/internal/application/job"
	"github.com/robfig/cron/v3"
	"os"
)

func main() {

	c := cron.New()
	job.JobCheckUpdateImage(c)
	c.Start()
	fmt.Println("⏱️ start cron jobs...")

	if os.Getenv("POS_TYPE") == "MASTER" {
		router.SetupWebRouter()
	}
}
