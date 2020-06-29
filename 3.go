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
		select {
		case data := <-chans:
			log.Printf("Row %v : %v -> done\n", data.indentity, data.content)
			wg.Done()
		case x := <-isClose:
			if x {
				fmt.Println("close")
				break
			}
		}
	}
}

func goRoutine3() {
	buffReadData := make(chan *DataLine, 10)
	defer close(buffReadData)
	var wg sync.WaitGroup
	isClose := make(map[int](chan bool), 3)

	f, _ := os.Open("file.txt")
	defer f.Close()

	for i := 0; i < 3; i++ {
		isClose[i] = make(chan bool, 1)
		go printData(buffReadData, &wg, isClose[i])
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

	for _, ch := range isClose {
		ch <- true
	}
}
