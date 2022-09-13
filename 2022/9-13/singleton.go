/******************************************************************************
* FILENAME:      singleton.go
*
* AUTHORS:       Xie Rongwang START DATE: 周二 9月 13 2022
*
* LAST MODIFIED: 星期二, 九月 13th 2022, 上午10:03
*
* CONTACT:       rongwang.xie@smartmore.com
******************************************************************************/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

type singleton struct {
	random int32
}

func main() {
	for i := 0; i < 10; i++ {
		singleton := Instance()
		fmt.Println(singleton.random)
	}
}

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &singleton{random: rand.Int31n(100)}
		defer atomic.StoreUint32(&initialized, 1)
	}
	return instance
}
