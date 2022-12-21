package main

import (
	"fmt"
	"os"
)

func createFile(file string) *os.File {
	f, err := os.Create(file)
	if err != nil {
		fmt.Println("Error occured")
		os.Exit(1)
	}
	return f
}


func main() {
	f := createFile("hello.txt")
        defer f.Close()
        fmt.Fprintln(f,"Hello World")
        fmt.Fprintln(f,"Bye World")
}
