package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter the token:")
		scanner.Scan()
		token := scanner.Text()
		switch token {
		case "A":
			fmt.Println("Your sit is on first row")
			break
		case "B":
			fmt.Println("Your sit is on second row")
			break
		case "C":
			fmt.Println("Your sit is on third row")
			break
		default:
			fmt.Println("Get your ticket checked")
		}
	}
}
