package main

import "fmt"

func main() {
	students := make(map[string]map[string]string)
	students["01"] = make(map[string]string)
	students["01"]["name"] = "小明"
	students["01"]["age"] = "18"
	students["02"] = make(map[string]string)
	students["02"]["name"] = "小话"
	students["02"]["age"] = "19"

	for k, v := range students {
		for k2, v2 := range v {
			fmt.Printf("students[%v][%v] = %v", k, k2, v2)
		}
		fmt.Println()
	}
}
