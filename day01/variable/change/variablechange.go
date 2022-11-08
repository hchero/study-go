/*
字符转换

*/

package main

import (
	"fmt"
)

func main() {
	/*
		转换规则 T(v) ，将值v转换为类型T
		T:就是数据类型，比如int32,int64,float32
		v:就是需要转换的变量

		范围大的转换为范围小的，会溢出

	*/
	var i int32 = 100
	// 将i转换为float32
	// var n1 float32 = i
	var n1 float32 = float32(1)
	fmt.Printf("i=%v,n1=%v", i, n1)

	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Printf("num1=%d,num2=%d", num1, num2)

}
