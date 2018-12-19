package main

import "fmt"

func shuffle(str string, char string) []string {
	res := make([]string, 0)

	// 递归基（终止条件）
	if str == "" {
		res = append(res, char)
		return res
	}

	// 取当前字符以及后面的字符串
	newChar := str[0:1]
	newStr := str[1:]

	strs := shuffle(newStr, newChar)
	for _, ranStr := range strs {
		res = append(res, char+ranStr)
		for i := 0; i < len(ranStr); i++ {
			res = append(res, ranStr[0:i+1]+char+ranStr[i+1:])
		}
	}

	return res
}

func main() {
	//str := "bc"
	//res := shuffle(str, "a")

	str := "23"
	res := shuffle(str, "1")

	for _, ranStr := range res {
		fmt.Println(ranStr)
	}
}
