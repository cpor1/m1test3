package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

type DataLine struct {
	content   string
	indentity int
}

func printData(chans chan *DataLine, wg *sync.WaitGroup, isClose chan bool) {

	for {
		data, j := <-chans
		if j {
			log.Printf("Row %v : %v -> done\n", data.indentity, data.content)
			wg.Done()
		} else {
			fmt.Println("Closed")
			isClose <- true
			return
		}
	}
}

func goRoutine3() {
	buffReadData := make(chan *DataLine, 10)
	defer close(buffReadData)
	var wg sync.WaitGroup
	isClose := make(chan bool)

	f, _ := os.Open("file.txt")
	defer f.Close()

	for i := 1; i <= 3; i++ {
		go printData(buffReadData, &wg, isClose)
	}
	scanner := bufio.NewScanner(f)
	count := 1
	for scanner.Scan() {
		dataLine := &DataLine{content: scanner.Text(), indentity: count}
		count++
		buffReadData <- dataLine
		wg.Add(1)
	}
	wg.Wait()
}
