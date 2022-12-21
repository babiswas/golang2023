package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter your age")
		scanner.Scan()
		age, _ := strconv.Atoi(scanner.Text())
		if age < 16 {
			fmt.Println("You are too young to join the club:")
		} else if age > 16 && age < 50 {
			fmt.Println("Welcome to the club")
		} else {
			fmt.Println("You are too old to join the club:")
		}
	}
}
