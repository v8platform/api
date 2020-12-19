package v8

import (
	"github.com/v8platform/enterprise"
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	type args struct {
		file   string
		params []map[string]string
	}
	tests := []struct {
		name string
		args args
		want enterprise.ExecuteOptions
	}{
		{
			"run epf",
			args{
				file:   "./path_to_epf.epf",
				params: nil,
			},
			enterprise.ExecuteOptions{
				File:   "./path_to_epf.epf",
				Params: nil,
			},
		},
		{
			"run epf with params",
			args{
				file: "./path_to_epf.epf",
				params: []map[string]string{
					{
						"param1": "any_string",
						"param2": "any_string",
					},
				},
			},
			enterprise.ExecuteOptions{
				File: "./path_to_epf.epf",
				Params: map[string]string{
					"param1": "any_string",
					"param2": "any_string",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Execute(tt.args.file, tt.args.params...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
