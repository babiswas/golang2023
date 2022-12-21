package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	fmt.Println("The length of the slice is:", len(arr1))
	fmt.Println("The capacity of the slice is:", cap(arr1))
	arr2 := make([]int, 5, 15)
	for i := 0; i < 5; i++ {
		arr2[i] = i + 10
	}
	fmt.Println("The length of the slice is:", len(arr2))
	fmt.Println("The capacity of the slice is:", cap(arr2))
	for index, value := range arr2 {
		fmt.Println(index, value)
	}
}
