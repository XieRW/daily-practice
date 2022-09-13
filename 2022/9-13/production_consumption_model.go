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
	"time"
)

func Producer(factor int, in chan<- int) {
	for i := 0; ; i++ {
		in <- i * factor
	}
}

func Consumer(out <-chan int) {
	for o := range out {
		fmt.Println(o)
	}
}

func main() {
	ch := make(chan int, 64)

	go Producer(2, ch)
	go Producer(3, ch)
	go Consumer(ch)

	time.Sleep(5 * time.Duration(50000))
}
