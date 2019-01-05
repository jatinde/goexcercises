package main

import (
	"fmt"
	"time"
)

func selectChannel(wordChannel chan string, doneChannel chan bool) {

	wordSlice := make([]string, 0, 4)
	wordSlice = append(wordSlice, "The")
	wordSlice = append(wordSlice, "Quick")
	wordSlice = append(wordSlice, "Brown")
	wordSlice = append(wordSlice, "Fox")
	count := 0
	t := time.NewTimer(3 * time.Second)
	defer close(wordChannel)
	defer close(doneChannel)
	for {
		select {
		case wordChannel <- wordSlice[count]:
			count = count + 1
			if count == len(wordSlice) {
				count = 0
			}
		case <-doneChannel:
			return
		case <-t.C:
			return
		}

	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go selectChannel(wordCh, doneCh)

	for word := range wordCh {
		fmt.Printf("%s ", word)
	}

}
