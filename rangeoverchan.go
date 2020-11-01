package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	//close(queue)

	go func() {
		close(queue)
	}()
	// 不能range没有关闭的chanel
	for elem := range queue {
		fmt.Println(elem)
	}

}
