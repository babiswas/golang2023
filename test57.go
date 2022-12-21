package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	prefix := make(chan string, 10)
	long_prefix := make(chan string, 10)
	var wg sync.WaitGroup
	wg.Add(1)

	go func(file_name string, wg *sync.WaitGroup, ch chan<- string) {
		defer wg.Done()
		file, err := os.Open(file_name)
		if err != nil {
			fmt.Println("Error occured:")
			os.Exit(1)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}("prefix.txt", &wg, prefix)

	max := -999999999
	for {
		select {
		case msg, ok := <-prefix:
			if ok {
				long_prefix <- msg
				for value := range prefix {
					long_prefix <- value
				}
			} else {
				close(long_prefix)
			}
		case val, ok := <-long_prefix:
			if ok {
				if max < len(val) {
					max = len(val)
				}
				for value := range long_prefix {
					if len(value) > max {
						max = len(value)
					}

				}
			} else {
				break
			}
		}
	}
	wg.Wait()
	fmt.Println("The maximum length is:", max)
}
