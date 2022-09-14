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
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("hello, world!")
			wg.Done()
		}()
	}

	wg.Wait()
}
