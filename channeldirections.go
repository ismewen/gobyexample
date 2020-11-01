package main

import "fmt"

// 把chan 换成 信道， 保持 <- 看数据的流向，流出为只读，流入为只写
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passmessage")
	fmt.Println(pings)
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
