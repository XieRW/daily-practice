/******************************************************************************
* FILENAME:      goroutine_helloworld.go
*
* AUTHORS:       Xie Rongwang START DATE: 周三 9月 14 2022
*
* LAST MODIFIED: 星期三, 九月 14th 2022, 下午5:15
*
* CONTACT:       rongwang.xie@smartmore.com
******************************************************************************/

package main

import (
	"fmt"
)

func main() {
	done := make(chan int, 10)

	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("hello, world!")
			done <- 1
		}()
	}

	for i := 0; i < cap(done); i++ {
		<-done
	}
}
