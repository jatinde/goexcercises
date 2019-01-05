package main

import "fmt"

func emitNumbers(c chan int64) {
	counter := int64(0)
	defer close(c)
	for counter < 100 {
		counter = counter + 1
		c <- counter
	}
}

func main() {
	numsChannel := make(chan int64)
	go emitNumbers(numsChannel)

	for num := range numsChannel {
		fmt.Printf("%d\n", num)
	}
}
