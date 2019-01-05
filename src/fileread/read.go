package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./file.txt")
	defer file.Close()
	if err != nil {
		fmt.Printf(err.Error())
	}

	b := make([]byte, 100)
	_, err = file.Read(b)

	fmt.Printf("%s", string(b))
	file.Write([]byte("Foo bar"))
}
