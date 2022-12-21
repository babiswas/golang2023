package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	id   int
}

func main() {
	mynums := []int{12, 11, 10, 2, 20, 25, 8, 0}
	mystring := []string{"hello", "bello", "mello", "aello"}
	person := [5]Person{{"Bapan", 12}, {"Madan", 1000}, {"Tapan", 121}, {"Suhas", 23}, {"Rahul", 4}}
	for index, value := range person {
		fmt.Println(index, value)
	}
	sort.Ints(mynums)
	fmt.Println("After sorting:")
	fmt.Println(mynums)
	fmt.Println("Sorting slice of strings in golang:")
	sort.Strings(mystring)
	fmt.Println(mystring)
	fmt.Println(person)
	sort.Slice(person[:], func(i, j int) bool { return person[i].id < person[j].id })
	fmt.Println(person)
}
