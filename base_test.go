package v8

import (
	"reflect"
	"testing"
)

func TestParseConnectionString(t *testing.T) {
	type args struct {
		connectingString string
	}
	tests := []struct {
		name    string
		args    args
		wantIb  *Infobase
		wantErr bool
	}{
		{
			"file",
			args{"File=./file_ib;Locale=ru_RU"},
			&Infobase{
				Connect: FilePath{File: "./file_ib"},
				Locale:  "ru_RU",
			},
			false,
		},
		{
			"server",
			args{"Srvr=test_server;Ref=test_ib;"},
			&Infobase{
				Connect: ServerPath{
					Server: "test_server",
					Ref:    "test_ib",
				},
			},
			false,
		},

		{
			"no file or server",
			args{"FFF=./file_ib;Locale=ru_RU"},
			nil,
			true,
		},

		{
			"ignore other values",
			args{"/UC 112;/UseTemplate;File=./file_ib;Locale=ru_RU"},
			&Infobase{
				Connect: FilePath{File: "./file_ib"},
				Locale:  "ru_RU",
			},
			false,
		},
		{
			"full server",
			args{"Srvr=test_server;Ref=test_ib;Usr=User;Pwd=Password;Prmod=1;LicDstr=Y"},
			&Infobase{
				Connect: ServerPath{
					Server: "test_server",
					Ref:    "test_ib",
				},
				User:                "User",
				Password:            "Password",
				AllowServerLicenses: true,
				UsePrivilegedMode:   true,
			},
			false,
		},

		{
			"full file",
			args{"File=./file_ib;Locale=ru_RU;Usr=User;Pwd=Password;Prmod=1;LicDstr=Y"},
			&Infobase{
				Connect: FilePath{
					File: "./file_ib",
				},
				User:                "User",
				Password:            "Password",
				AllowServerLicenses: true,
				UsePrivilegedMode:   true,
				Locale:              "ru_RU",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIb, err := ParseConnectionString(tt.args.connectingString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConnectionString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIb, tt.wantIb) {
				t.Errorf("ParseConnectionString() gotIb = %v, want %v", gotIb, tt.wantIb)
			}
		})
	}
}

func TestInfobase_ConnectionString(t *testing.T) {
	type fields struct {
		Connect             ConnectPath
		User                string
		Password            string
		AllowServerLicenses bool
		SeparatorList       DatabaseSeparatorList
		UsePrivilegedMode   bool
		Locale              string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"simple file",
			fields{
				Connect: FilePath{
					File: "./test_ib",
				},
				User:     "user",
				Password: "password",
			},
			"/IBConnectionString File='./test_ib';Usr=user;Pwd=password",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ib := Infobase{
				Connect:             tt.fields.Connect,
				User:                tt.fields.User,
				Password:            tt.fields.Password,
				AllowServerLicenses: tt.fields.AllowServerLicenses,
				SeparatorList:       tt.fields.SeparatorList,
				UsePrivilegedMode:   tt.fields.UsePrivilegedMode,
				Locale:              tt.fields.Locale,
			}
			if got := ib.ConnectionString(); got != tt.want {
				t.Errorf("ConnectionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
