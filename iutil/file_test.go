package iutil

import (
	"testing"
)

type Config struct {
	Rmq RmqConfig `json:"rmq" yaml:"rmq"`
}

type RmqConfig struct {
	Consumer []ConsumerConf `json:"consumer" yaml:"consumer"`
}

type ConsumerConf struct {
	Service    string   `json:"service" yaml:"service"`
	NameServer string   `json:"nameserver" yaml:"nameserver"`
	Topic      string   `json:"topic" yaml:"topic"`
	Tags       []string `json:"tags" yaml:"tags"`
	Group      string   `json:"group" yaml:"group"`
	Broadcast  bool     `json:"broadcast" yaml:"broadcast"`
	Orderly    bool     `json:"orderly" yaml:"orderly"`
	Retry      int      `json:"retry" yaml:"retry"`
}

func TestGetRootPath(t *testing.T) {
	path := "/root"
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success",
			want: path,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRootPath(); got != DefaultRootPath {
				t.Errorf("GetRootPath() = %v, want %v", got, DefaultRootPath)
			}
			SetRootPath(path)
			if got := GetRootPath(); got != tt.want {
				t.Errorf("GetRootPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadYaml(t *testing.T) {
	type args struct {
		filename string
		subPath  string
		s        interface{}
	}
	var test Config
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				filename: "test.yaml",
				subPath:  "",
				s:        &test,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoadYaml(tt.args.filename, tt.args.subPath, tt.args.s)
			t.Log(ToJson(tt.args.s))
		})
	}
}
