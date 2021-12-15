package iutil

import jsoniter "github.com/json-iterator/go"

// ToJson 用了jsoniter 的MarshalToString，不抛出异常，慎用！！！
func ToJson(input interface{}) string {
	res, _ := jsoniter.MarshalToString(input)
	return res
}
