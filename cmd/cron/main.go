package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("cron started.")

	c := cron.New()
	c.AddFunc("*/1 * * * *", func() { fmt.Println("a new minute.") })
	c.Start()

	select {}
}
