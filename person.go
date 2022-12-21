package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	id   int
}

func (p *Person) setname(name string) {
	p.name = name
}

func (p *Person) setid(id int) {
	p.id = id
}

func (p Person) display() string {
	return "Name is:" + p.name + " " + "Id is:" + strconv.Itoa(p.id)
}

func main() {
	var p Person
	p.setname("Bapan")
	p.setid(12)
	fmt.Println(p)
	fmt.Println(p.display())
}
