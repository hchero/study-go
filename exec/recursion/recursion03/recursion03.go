package main

import "fmt"

// func hzct(n int) int {
// 	if n == 1 {
// 		fmt.Println(1)
// 		return 1
// 	} else {
// 		result := (hzct(n-1) + 1) * 2
// 		fmt.Println(result)
// 		return result
// 	}
// }

func hzct(n int) int {
	if n == 10 {
		return 1
	}
	if n > 10 || n < 1 {
		fmt.Println("天数不对")
		return 0
	} else {
		return (hzct(n+1) + 1) * 2
	}
}

/*
n =n/2 -1



n,(n+1) * 2
1,4,10,22
1,(1+1)*2,

*/

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(hzct(i))
	}
}
