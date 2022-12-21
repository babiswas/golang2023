package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func process_file(wg *sync.WaitGroup, filename string, ch chan<- string) {
	defer wg.Done()
	file, err := os.Open(filename)
	if err != nil {
		close(ch)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
	close(ch)
}

func process_string(wg *sync.WaitGroup, ch <-chan string, ch1 chan<- string, ref string) {
	msg := ""
	var test bool
	defer wg.Done()
	for {
		select {
		case msg, test = <-ch:
			if test == true {
				if strings.HasPrefix(ref, msg) {
					ch1 <- msg
				}
			}
		}
		if test == false {
			close(ch1)
			break
		}
	}
	return
}

func process_length(wg *sync.WaitGroup, ch1 <-chan string, ch3 chan<- int) {
	defer wg.Done()
	var test bool
	max := 0
	num := ""
	for {
		select {
		case num, test = <-ch1:
			if test == true {
				if len(num) > max {
					max = len(num)
				}
			}
		}
		if test == false {
			ch3 <- max
			break
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)
	ch3 := make(chan int, 1)
	scanner1 := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the input strings:")
	scanner1.Scan()
	ref := scanner1.Text()
	wg.Add(1)
	go process_file(&wg, "prefix.txt", ch1)
	wg.Add(1)
	go process_string(&wg, ch1, ch2, ref)
	wg.Add(1)
	go process_length(&wg, ch2, ch3)
	fmt.Println(<-ch3)
	wg.Wait()
}
