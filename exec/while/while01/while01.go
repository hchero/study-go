package main

import "fmt"

func main() {
	var totalScore float32
	var classScore float32
	for i := 1; i <= 3; i++ {
		classScore = 0
		fmt.Printf("班级%d\n", i)
		for j := 1; j <= 5; j++ {
			var score float32
			fmt.Printf("请输入第%d个学生的分数\n", j)
			fmt.Scanln(&score)
			classScore += score
		}
		totalScore += classScore
		avg := classScore / 5
		fmt.Printf("班级%d的平均分是%f\n", i, avg)
	}
	totalAvg := totalScore / 15
	fmt.Printf("所有班级的平均分是%f\n", totalAvg)
}
