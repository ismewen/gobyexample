package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zerovalptr(pval *int) {
	*pval = 0
}

func main(){
	i := 1
	fmt.Println("initial", i)
	zeroval(i)
	fmt.Println(i)
	zerovalptr(&i)
	fmt.Println(i)
}
