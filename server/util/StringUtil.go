package util

import (
	"fmt"
	"strconv"
	"strings"
)

const EMPTY string = ""

/**
 * 空字符串
 */
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

/**
 * 非空字符串
 */
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

/**
 * 首字母大写
 */
func FirstToUpperCase(str string) string {
	return strings.ToUpper(Substring(str, 0, 1)) + Substring(str, 1, len(str))
}

/**
 * 首字母小写
 */
func FirstToLowerCase(str string) string {
	return strings.ToLower(Substring(str, 0, 1)) + Substring(str, 1, len(str))
}

/**
 * 截取字符串
 */
func SubLastCharacter(str string) string {
	return Substring(str, 0, len(str)-1)
}

//获取source的子串,如果start小于0或者end大于source长度则返回""
//start:开始index，从0开始，包括0
//end:结束index，以end结束，但不包括end
func Substring(source string, start int, end int) string {
	r := []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func ArrayToString(arr []int) string {
	var result string
	//遍历数组中所有元素追加成string
	for i, value := range arr {
		if i != 0 {
			result = result + " "
		}
		result = result + ZeroFillByStr(IntToStr(value), 2, true)
	}
	return result
}

//
//  ZeroFillByStr
//  @Description: 字符串补零
//  @param str :需要操作的字符串
//  @param resultLen 结果字符串的长度
//  @param reverse true 为前置补零，false 为后置补零
//  @return string
//
func ZeroFillByStr(str string, resultLen int, reverse bool) string {
	if len(str) > resultLen || resultLen <= 0 {
		return str
	}
	if reverse {
		return fmt.Sprintf("%0*s", resultLen, str) //不足前置补零
	}
	result := str
	for i := 0; i < resultLen-len(str); i++ {
		result += "0"
	}
	return result
}
