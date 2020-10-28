package main

import (
	"fmt"
	"os"
)

func main(){
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set", s)
	fmt.Println("get", s[2])
	fmt.Println("len", len(s))

	b := append(s, "d")
	b = append(s, "e", "f")

	fmt.Println(b)
	fmt.Println(s)

	s = b

	c := make([]string, len(s))
	copy(c, s)

	l := s[2:5]

	fmt.Println("cpy", c)
	fmt.Println("sli", l)

}
