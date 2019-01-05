package main

import (
	"fmt"
	"time"
)

func printer(message string, msgCh chan bool) {
	for {
		select {
		case <-msgCh:
			return
		default:
			fmt.Printf("%s\n", message)
		}
	}
}

func main() {
	msgCh := make(chan bool)

	for i := 1; i <= 10; i++ {
		go printer(fmt.Sprintf("Printing %d", i), msgCh)
	}
	time.Sleep(5 * time.Second)
	close(msgCh)
	time.Sleep(5 * time.Second)
}
