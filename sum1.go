package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the second number:")
	scanner.Scan()
	num1, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Enter second number:")
	scanner.Scan()
	num2, _ := strconv.Atoi(scanner.Text())
	fmt.Println("The sum of the number is:", num1+num2)
}
