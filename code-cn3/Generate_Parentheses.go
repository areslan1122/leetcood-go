package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DataEditor(data []string) {

	tmpData := map[string]int64{} //保存去重后的数据
	tmpSArr := []string{} //临时字符串排序数组
	tmpTArr := []int64{} //临时时间戳排序

	//去掉重复字符串
	for _, str := range data {
		keyvalue := strings.Split(str, ":")

		//获取单个输入
		key, _ := strconv.ParseInt(keyvalue[0], 10, 64)
		value := keyvalue[1]

		if key > tmpData[value] {
			tmpData[value] = key
		}
	}


	for key, value := range tmpData {
		tmpSArr = append(tmpSArr, key)
		tmpTArr = append(tmpTArr, value)
	}


	tmpSArr = stringArrSort(tmpSArr)
	tmpTArr = intArrSort(tmpTArr)

	for i:=0; i<len(tmpTArr); i++ {
		for j:=0; j<len(tmpSArr); j++ {
			if tmpSArr[j] != "" {
				for key, value := range tmpData {
					if value == tmpTArr[i] && key == tmpSArr[j] {
						fmt.Println("timestamp:", value, " value:", key)
						tmpSArr[j] = ""
					}
				}
			}
		}
	}
}

//数字数组排序
func intArrSort(arr []int64) []int64 {

	for i:=0; i<len(arr); i++ {
		for j:=i+1; j<len(arr); j++ {
			if arr[j] < arr[i] {
				tmp := arr[i]
				arr[i] =arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr
}


//字符串数组排序
func stringArrSort(str []string) []string {

	for i:=0; i<len(str); i++ {
		for j:=i+1; j<len(str); j++ {
			if compareString(str[i], str[j]) == str[i] {
				continue
			} else {
				tmp := str[i]
				str[i] =str[j]
				str[j] = tmp
			}
		}
	}
	return str
}

//两个字符串前后顺序比较
func compareString(a,b string) string {
	var long int
	if len(a) >= len(b) {long = len(a)} else {long = len(b)}
	for i:=0; i<long; i++ {
		if a[i] > b[i] {
			return a
		} else if b[i] > a[i] {
			return b
		}
	}
	if long == len(a) {
		return b
	} else {
		return a
	}
}

func main() {

	data := []string{
		"1570593273486:03Ad2k33",
		"1570593273487:03Ad233d988dfd",
		"1570593273488:03Ad2k34",
		"1570593273488:03Ad233d988dfd",
		"1570593273489:03Ad2k33",
		"1570593273487:03Ad233d988dfd",
		"1570593273486:03Ad2k34",
		"1570593273492:03Ad233d988dfd",
		"1570593273493:03Ad2k33",
		"1570593273494:03Ad233d988dfd",
		"1570593273494:03Ad234d988dfd",
	}
	DataEditor(data)
}