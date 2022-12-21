package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	for _, value := range os.Args {
		num, _ := strconv.Atoi(value)
		sum += num
	}
	fmt.Println("The sum of the numbers is:", sum)
}
