package main

import "fmt"

func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = i + 100
	}
	for index, value := range arr {
		fmt.Println(index, value)
	}
}
