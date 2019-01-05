package main

import (
	"fmt"
	"os"
)

func main() {
	if file, err := os.Create("abc.txt"); err != nil {
		fmt.Println(err.Error())
	} else {
		defer file.Close()
		file.WriteString("Hey there from visual studio code again!")
	}
}
