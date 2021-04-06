package go_eudic

import (
	"encoding/json"
	"regexp"
)

func StructToString(v interface{}) (string, error) {
	marshal, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(marshal), err
}

// 获取字符串参数中的首个数字
func GetTheFirstNumberFromString(stringContainsNumber string) string {
	numberRegexp := regexp.MustCompile("[0-9]+")
	foundString := numberRegexp.FindString(stringContainsNumber)
	return foundString
}