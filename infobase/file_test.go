package infobase

import "testing"

func TestFile_Parse(t *testing.T) {
	type fields struct {
		Common Common
		File   string
		Locale string
	}
	type args struct {
		connectingString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Simple",
			fields{
				Common: Common{},
				File:   "./file_ib",
				Locale: "ru_RU",
			},
			args{
				connectingString: "",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ib := File{
				Common: tt.fields.Common,
				File:   tt.fields.File,
				Locale: tt.fields.Locale,
			}
			if err := ib.Parse(tt.args.connectingString); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
