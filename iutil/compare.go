package iutil

import "github.com/mozillazg/go-pinyin"

/*
HanLess

拼音相同，字不同，比较编码大小
字相同，比较下一个字
有字 > 无字
字符串完全相同返回ture
*/
func HanLess(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	s1Arr := []rune(s1)
	s2Arr := []rune(s2)
	minLen := MinInt(len(s1Arr), len(s2Arr))
	for i := 0; i < minLen; i++ {
		if s1Arr[i] != s2Arr[i] {
			pin1 := pinyin.LazyConvert(string(s1Arr[i]), nil)[0]
			pin2 := pinyin.LazyConvert(string(s2Arr[i]), nil)[0]
			if pin1 == pin2 {
				return s1Arr[i] < s2Arr[i]
			}
			return pin1 < pin2
		}
	}
	return len(s1Arr) < len(s2Arr)
}

// MinInt 返回较小值
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt 返回较大值
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
