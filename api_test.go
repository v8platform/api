package v8

import (
	"github.com/v8platform/designer"
	"github.com/v8platform/runner"
	"reflect"
	"testing"
)

func TestCreateInfobase(t *testing.T) {

	if testing.Short() {
		t.Skip("skipped for integrated tests")
	}

	type args struct {
		create runner.Command
		opts   []interface{}
	}
	temp := t.TempDir()

	tests := []struct {
		name    string
		args    args
		want    *Infobase
		wantErr bool
	}{
		{
			"simple",
			args{
				create: designer.CreateFileInfoBaseOptions{
					File: temp,
				},
			},
			&Infobase{
				Connect: FilePath{
					File: temp,
				},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateInfobase(tt.args.create, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateInfobase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInfobase() got = %v, want %v", got, tt.want)
			}
		})
	}
}
