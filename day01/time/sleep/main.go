package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		i++
		fmt.Println(i)
		if i >= 10 {
			break
		}
		time.Sleep(time.Microsecond * 100)
	}

}
