package iutil

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const DefaultRootPath = "."

var rootPath string

// SetRootPath 设置应用的根目录
func SetRootPath(r string) {
	rootPath = r
}

// RootPath 返回应用的根目录
func GetRootPath() string {
	if rootPath != "" {
		return rootPath
	} else {
		return DefaultRootPath
	}
}

// LoadYaml 加载yaml 文件内容到结构体
func LoadYaml(filename, subPath string, s interface{}) {
	var path string
	path = filepath.Join(GetRootPath(), subPath, filename)

	if yamlFile, err := os.ReadFile(path); err != nil {
		panic(filename + " get error: %v " + err.Error())
	} else if err = yaml.Unmarshal(yamlFile, s); err != nil {
		panic(filename + " unmarshal error: %v" + err.Error())
	}
}
