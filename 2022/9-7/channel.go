/******************************************************************************
* FILENAME:      channel.go
*
* AUTHORS:       Xie Rongwang START DATE: 周三 9月 07 2022
*
* LAST MODIFIED: 星期三, 九月 07th 2022, 下午4:32
*
* CONTACT:       rongwang.xie@smartmore.com
******************************************************************************/

package main

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好, 世界"
	close(done)
}

func main() {
	go aGoroutine()
	<-done
	println(msg)
}
