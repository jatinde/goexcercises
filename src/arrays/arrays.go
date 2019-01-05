package main

import "fmt"

func main() {
	nums := [...]int64{1, 3, 5, 7, 6, 4, 2}
	arrayPrinter(nums)
	arrayPrinter(nums)
}

func arrayPrinter(nums [7]int64) {
	for i, num := range nums {
		fmt.Printf("%d: %d\n", i, num)
	}
	nums[2] = 9
}
