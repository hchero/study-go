package main

import "fmt"

func main() {
	var studentCount int
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			var score float32
			fmt.Printf("请输入班级(%d)第%d个学生的分数\n", i, j)
			fmt.Scanln(&score)
			if score >= 60 {
				studentCount++
			}
		}
	}
	fmt.Printf("三个班及格人数：%d\n", studentCount)
}
