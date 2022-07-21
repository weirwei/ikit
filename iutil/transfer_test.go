package iutil

import "testing"

type A struct {
	Str     string `json:"str"`
	Integer int    `json:"integer"`
	Obj     struct {
		Str     string `json:"str"`
		Integer int    `json:"integer"`
	} `json:"obj"`
}

func TestStructMap(t *testing.T) {
	type args struct {
		st interface{}
		m  map[string]interface{}
	}
	m := make(map[string]interface{})
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				st: A{
					Str:     "123",
					Integer: 1,
					Obj: struct {
						Str     string `json:"str"`
						Integer int    `json:"integer"`
					}{
						Str:     "456",
						Integer: 7,
					},
				},
				m: m,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StructMap(tt.args.st, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("StructMap() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(ToJson(tt.args.m))
		})
	}
}
