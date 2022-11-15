package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

// 实现学生的切片
type StudentSlice []Student

func (ss StudentSlice) Len() int {
	return len(ss)
}

func (ss StudentSlice) Less(i, j int) bool {
	if ss[i].Age > ss[j].Age {
		return false
	} else {
		return true
	}
}

func (ss StudentSlice) Swap(i, j int) {
	temp := ss[j]
	ss[j] = ss[i]
	ss[i] = temp
}

func main() {
	var students StudentSlice
	students = append(students, Student{Name: "小明", Age: 17, Score: 90})
	students = append(students, Student{Name: "小王", Age: 16, Score: 100})
	students = append(students, Student{Name: "小李", Age: 19, Score: 99})
	sort.Sort(students)
	fmt.Println(students)
}
