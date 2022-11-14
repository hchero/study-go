/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-14 16:36:01
 * @LastEditors: super
 * @Description:
 */
package model

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func NewStudent(name string, age int, score int) *Student {
	return &Student{
		Name:  name,
		Age:   age,
		Score: score,
	}
}

type Pupil struct {
	Student // 嵌入Student
}

// 大学生 college

type College struct {
	Student // 嵌入Student
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

func (stu *Student) GetScore() int {
	return stu.Score
}

func (pupil *Pupil) Testing() {
	fmt.Println("小学生正在考试..")
}

func (college *College) Testing() {
	fmt.Println("大学生正在考试")
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生姓名：%v,年龄：%v,分数：%v", stu.Name, stu.Age, stu.Score)
}
