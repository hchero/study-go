package main

import "fmt"

/*
输入五个学生成绩并打印
*/
func main() {
	var score [5]int
	for i := 0; i < 5; i++ {
		fmt.Printf("输入第%v个学生成绩\n", i+1)
		fmt.Scanln(&score[i])
	}

	for i := 0; i < len(score); i++ {
		fmt.Println(score[i])
	}
}
