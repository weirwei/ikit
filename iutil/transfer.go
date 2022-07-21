package iutil

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

// 这几个byte 相关的是抄的 https://studygolang.com/articles/11981

// StringBytes return GoString's buffer slice(enable modify string)
func StringBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// BytesString convert b to string without copy
func BytesString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StructMap transfer struct to map[string]interface{}
func StructMap(st interface{}, m map[string]interface{}) error {
	bytes, err := jsoniter.Marshal(st)
	if err != nil {
		return err
	}
	if err = jsoniter.Unmarshal(bytes, &m); err != nil {
		return err
	}
	return nil
}
