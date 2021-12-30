package iutil

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// LoadYaml 加载yaml 文件内容到结构体
func LoadYaml(filename, subPath string, s interface{}) {
	var path string
	path = filepath.Join(subPath, filename)

	if yamlFile, err := os.ReadFile(path); err != nil {
		panic(filename + " get error: %v " + err.Error())
	} else if err = yaml.Unmarshal(yamlFile, s); err != nil {
		panic(filename + " unmarshal error: %v" + err.Error())
	}
}
