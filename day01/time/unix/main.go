package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("unix时间戳=%v，UnixNano时间戳=%v", time.Now().Unix(), time.Now().UnixNano())
}
