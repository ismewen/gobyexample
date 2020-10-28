package main

import "fmt"

func plus(a int, b int) (int, error) {
	return a + b, nil
}

func Sum(a ...int) int {
	var sum int
	fmt.Println("here", a)
	for x, v := range a {
		fmt.Println(x, v)
		sum += v
	}
	return sum
}

func main() {
	sum, _ := plus(1, 2)
	fmt.Println(sum)
	s := Sum(1, 2, 3, 4)
	fmt.Println(s)

}
