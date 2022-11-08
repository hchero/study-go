package main

import "fmt"

func main() {
	var total int
	var current int
	i := 1
	for {
		if current == 0 && total > 20 {
			current = i
		}
		total += i
		if i >= 100 {
			break
		}
		i++
	}
	fmt.Printf("100以内的数的和是%v,第一次大于100的当前数是%v", total, current)
}
