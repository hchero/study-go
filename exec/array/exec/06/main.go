/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-08 10:23:40
 * @LastEditors: super
 * @LastEditTime: 2022-11-08 11:18:21
 * @Description:试写出实现查找的核心代码，比如已知数组arr[10]string里面保存了10个元素，现要查找"AA"在其中是否存在，打印提示，如果有多个"AA"，也要找到对应的下标
 */
package main

import "fmt"

func find(arr *[]string, findValue string) (bool, []int) {
	var indexList []int
	var isFind bool
	for i := 0; i <= len(*arr)/2; i++ {
		if (*arr)[i] == findValue {
			isFind = true
			indexList = append(indexList, i)
			// break
		}
		if (*arr)[len(*arr)-1-i] == findValue {
			isFind = true
			indexList = append(indexList, len(*arr)-1-i)
			// break
		}
	}
	return isFind, indexList
}

func find2(arr []string, findValue string) (bool, []int) {
	var indexList []int
	var isFind bool
	for i := 0; i <= len(arr)/2; i++ {
		if arr[i] == findValue {
			isFind = true
			indexList = append(indexList, i)
			// break
		}
		if arr[len(arr)-1-i] == findValue {
			isFind = true
			indexList = append(indexList, len(arr)-1-i)
			// break
		}
	}
	return isFind, indexList
}

func main() {
	var arr []string = []string{"AA", "bb", "CC", "bb", "EE", "FF", "DD", "HH", "AA", "bb"}
	// isHas, indexList := find(&arr, "AA")
	isHas, indexList := find2(arr, "AA")
	if isHas {
		fmt.Println("找到了，下标为")
		for i := 0; i < len(indexList); i++ {
			fmt.Printf("%v\t", indexList[i])
		}
	} else {
		fmt.Println("没找到")
	}

}
