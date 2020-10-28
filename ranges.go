package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	kvx := map[string]string{"a": "hello", "b": "world"}

	for key, value := range kvx{
		fmt.Println("key--->", key, "values--->", value)
	}

	for k := range kvx {
		fmt.Println("key-->", k)
	}

	for i, c := range "go"{
		fmt.Println(string(i), c, "go"[0])
	}

}
