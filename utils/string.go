package utils

import (
	"regexp"
	"strconv"
)

func ParseInt(str string) int {
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(res)
}

func GetRegex(str string) *regexp.Regexp {
	regex, err := regexp.Compile(str)
	if err != nil {
		panic(err)
	}
	return regex
}