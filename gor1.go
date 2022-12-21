package main

import (
	"fmt"
	"sync"
	"time"
)

func task(wg *sync.WaitGroup, str1 string) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("Executing task:", i, "Task:", str1)
		time.Sleep(3 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go task(&wg, "Mop floor")
	wg.Add(1)
	go task(&wg, "Clean Dish")
	wg.Wait()
}
