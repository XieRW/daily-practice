/******************************************************************************
* FILENAME:      production_consumption_model.go
*
* AUTHORS:       Xie Rongwang START DATE: 周二 9月 13 2022
*
* LAST MODIFIED: 星期二, 九月 13th 2022, 下午6:50
*
* CONTACT:       rongwang.xie@smartmore.com
******************************************************************************/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func producer(factor int, in chan<- int) {
	for i := 0; ; i++ {
		in <- i * factor
	}
}

func consumer(out <-chan int) {
	for o := range out {
		fmt.Println(o)
	}
}

func main() {
	ch := make(chan int, 64)

	go producer(2, ch)
	go producer(3, ch)
	go consumer(ch)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
