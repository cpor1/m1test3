package main

import (
	"log"
	"time"
)

func chanRoutine() {
	done := make(chan bool)
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		done <- true
	}()
	log.Print("hello 2")
	<-done
}
