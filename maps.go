package main

import "fmt"

func init() {
	fmt.Println("init one")
}
func init()  {
	fmt.Println("init two")

}
func main() {
	m := make(map[string]interface{})
	m["k1"] = 1
	m["k2"] = "hello"

	v1 := m["k1"]

	fmt.Println("v1", v1)
	fmt.Println("v2", m["k2"])

}
