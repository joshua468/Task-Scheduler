package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	_, err := c.AddFunc("*/1 * * * *", func() {
		fmt.Println("Running scheduled task:", time.Now())
	})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	_, err = c.AddFunc("0 2 * * *", func() {
		fmt.Println("Running daily task", time.Now())
	})
	if err != nil {
		fmt.Println("Error", err)
	}

	c.Start()

	<-time.After(5 * time.Minute)
	c.Stop()
}
