/******************************************************************************
* FILENAME:      singleton.go
*
* AUTHORS:       Xie Rongwang START DATE: 周五 9月 09 2022
*
* LAST MODIFIED: 星期五, 九月 09th 2022, 上午9:52
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
	instance             *singleton
	instanceNotSingleton *singleton
	initialized          uint32
	mu                   sync.Mutex
)

type singleton struct {
	random int64
}

func main() {
	for i := 0; i < 10; i++ {
		object := Instance()
		fmt.Println(object.random)
	}

	fmt.Println("=======================")

	for i := 0; i < 10; i++ {
		object := InstanceNotSingleton()
		fmt.Println(object.random)
	}
}

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{random: rand.Int63n(100)}
	}
	return instance
}

func InstanceNotSingleton() *singleton {
	instanceNotSingleton = &singleton{random: rand.Int63n(100)}
	return instanceNotSingleton
}
