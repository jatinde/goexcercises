package main

import "fmt"

func main() {
	nums := make([]int64, 0, 5)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	nums = append(nums, 1)
	nums = append(nums, 3)
	nums = append(nums, 5)
	nums = append(nums, 6)
	nums = append(nums, 2)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	nums = append(nums, 4)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	slicePrinter(nums[1:2])
}

func slicePrinter(nums []int64) {
	for i, num := range nums {
		fmt.Printf("%d: %d\n", i, num)
	}
}
