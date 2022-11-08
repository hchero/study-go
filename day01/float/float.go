package main
import (
	"unsafe"
	"fmt"
)
func main(){
	var num float32 = .1234;
	fmt.Println(num);
	fmt.Println("数字num的长度 %v",unsafe.Sizeof(num));
}