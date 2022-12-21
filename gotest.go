package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	count := 0
	prefix := make(chan string, 10)
	file, err := os.Open("prefix.txt")
	if err != nil {
		fmt.Println("Error occured:")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wg.Add(1)
		chr := scanner.Text()
		if len(chr) != 0 {
			go func(wg *sync.WaitGroup, ch chan<- string, str1 string) {
				defer wg.Done()
				ch <- str1
			}(&wg, prefix, chr)
			count += 1
		} else {
			close(prefix)
		}
	}
	for {
		if len(prefix) == 0 {
			close(prefix)
		}
		select {
		case msg, ok := <-prefix:
			if ok == false {
				break
			} else {
				if count != 0 {
					fmt.Println(msg)
					count -= 1
				}
				if count == 0 {
					close(prefix)
					break
				} else {
					for value := range prefix {
						fmt.Println(value)
						count -= 1
					}
				}
			}
		}
	}
}
