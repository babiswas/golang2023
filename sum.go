package main

import "fmt"

func add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := add(nums...)
	fmt.Println("The sum of the numbers is:", sum)
}
