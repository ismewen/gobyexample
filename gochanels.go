package main

import (
	"fmt"
	"time"
)

func main() {
	// 无缓冲的信道, 接收方要先发送方准备好，否则发送方就会阻塞
	message := make(chan string)
	go func() { message <- "ping" }()

	go func() {
		msg := <-message
		fmt.Println(msg)
	}()


	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "hello"
	//fmt.Println(<-messages)
	//fmt.Println(<-messages)

	go func() {
		fmt.Println(<-messages)
		fmt.Println(<-messages)
	}()

	time.Sleep(time.Second * 3)


}
