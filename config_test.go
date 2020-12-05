package v8

import (
	"github.com/stretchr/testify/assert"
	"github.com/v8platform/designer"
	"github.com/v8platform/runner"
	"reflect"
	"testing"
)

func TestLoadCfg(t *testing.T) {

	ib := NewFileInfobase("./test_ib")

	type args struct {
		file string
	}
	tests := []struct {
		name      string
		args      args
		want      designer.LoadCfgOptions
		want_args []string
	}{
		{
			"simple",
			args{file: "./1cv8.cf"},
			designer.LoadCfgOptions{File: "./1cv8.cf", Designer: designer.NewDesigner()},
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/LoadCfg ./1cv8.cf",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadCfg(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadCfg() = %v, want %v", got, tt.want)
			}

			v8run := runner.NewPlatformRunner(ib, LoadCfg(tt.args.file))
			got := v8run.Args()

			for _, arg := range tt.want_args {
				assert.Contains(t, got, arg,
					"NewPlatformRunner() = %v, want %v", got, arg)

			}

		})
	}
}
