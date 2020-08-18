package util

import "strconv"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseInt(str string) int {
	var result int
	var err error
	if IsNotBlank(str) {
		result, err = strconv.Atoi(str)
		CheckError(err)
	}
	return result
}