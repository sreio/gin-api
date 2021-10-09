package main

import (
	"gin-api/models"
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	log.Println("strating...")

	c := cron.New()
	c.AddFunc("* * * * * *", func ()  {
		log.Println("run models.CleanAllTag...")
		models.CleanAllTag()
	})

	c.AddFunc("* * * * * *", func ()  {
		log.Println("run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)

	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}