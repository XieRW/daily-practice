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
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("hello, world!")
		mu.Unlock()
	}()

	mu.Lock()
}
