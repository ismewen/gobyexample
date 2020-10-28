package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name, age: 42}
	return &p
}

func (p *Person) sayHello(){
	fmt.Println(p.name,"say hello")
}

func main(){
	ethan := newPerson("ethan")
	ethan.sayHello()
}