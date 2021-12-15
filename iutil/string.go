package iutil

import "regexp"

// Trim 去除所有的空白字符
func Trim(str string) string {
	if len(str) == 0 {
		return ""
	}
	return regexp.MustCompile(`\\s+`).ReplaceAllString(str, "")
}
