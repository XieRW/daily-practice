/******************************************************************************
* FILENAME:      sqint.go
*
* AUTHORS:       Xie Rongwang START DATE: 周四 7月 14 2022
*
* LAST MODIFIED: 星期四, 七月 14th 2022, 下午3:37
*
* CONTACT:       rongwang.xie@smartmore.com
******************************************************************************/

package main

import "fmt"

func main() {
	i := 10
	sq := sqint(i)
	fmt.Println(sq)
}

func sqint(i int) int {
	return i * i
}
