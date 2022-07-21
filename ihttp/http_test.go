package ihttp

import (
	"testing"

	"github.com/weirwei/ikit/iutil"
)

func TestGET(t *testing.T) {
	type args struct {
		opt *Options
	}
	tests := []struct {
		name    string
		args    args
		want    *Result
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				opt: &Options{
					URL: "https://suggest.taobao.com/sug",
					RequestBody: map[string]interface{}{
						"code": "utf-8",
						"q":    "ps5",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "retry",
			args: args{
				opt: &Options{
					URL: "https://suggest.taoxxxxbao.com/suxxxxxg",
					RequestBody: map[string]interface{}{
						"code": "utf-8",
						"q":    "ps5",
					},
					Retry: 2,
				},
			},
			wantErr: true,
		},
		{
			name: "timeout",
			args: args{
				opt: &Options{
					URL: "https://suggest.taobao.com/sug",
					RequestBody: map[string]interface{}{
						"code": "utf-8",
						"q":    "ps5",
					},
					Timeout: 3,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GET(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("GET() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(iutil.ToJson(got))
		})
	}
}

func TestPOST(t *testing.T) {
	type args struct {
		opt *Options
	}
	tests := []struct {
		name    string
		args    args
		want    *Result
		wantErr bool
	}{
		{
			name: "form success",
			args: args{
				opt: &Options{
					URL: "https://e.juejin.cn/resources/github",
					RequestBody: map[string]interface{}{
						"category": "treading",
						"period":   "day",
						"lang":     "go",
						"offset":   0,
						"limit":    2,
					},
					Encode: EncodeForm,
				},
			},
			wantErr: false,
		},
		{
			name: "json success",
			args: args{
				opt: &Options{
					URL: "https://e.juejin.cn/resources/github",
					RequestBody: map[string]interface{}{
						"category": "treading",
						"period":   "day",
						"lang":     "go",
						"offset":   0,
						"limit":    2,
					},
					Encode: EncodeJson,
				},
			},
			wantErr: false,
		},
		{
			name: "retry",
			args: args{
				opt: &Options{
					URL: "https://e.juejssssin.cn/resources/github",
					RequestBody: map[string]interface{}{
						"category": "treading",
						"period":   "day",
						"lang":     "go",
						"offset":   0,
						"limit":    2,
					},
					Encode: EncodeJson,
					Retry:  3,
				},
			},
			wantErr: true,
		},
		{
			name: "timeout",
			args: args{
				opt: &Options{
					URL: "https://e.juejin.cn/resources/github",
					RequestBody: map[string]interface{}{
						"category": "treading",
						"period":   "day",
						"lang":     "go",
						"offset":   0,
						"limit":    2,
					},
					Encode:  EncodeJson,
					Timeout: 3,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := POST(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("POST() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(iutil.ToJson(got))
		})
	}
}
