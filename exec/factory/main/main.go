package main

import (
	"fmt"
	"go_code/project01/exec/factory/module"
)

func main() {
	var stu = module.NewStudent("小明", 99.5)
	fmt.Println(stu)
}
