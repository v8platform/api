package v8

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"os"
	"reflect"
	"testing"
)

type baseTestSuite struct {
	suite.Suite
}

func (b *baseTestSuite) SetupSuite() {

}

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func TestFileInfobaseFromConnectingString(t *testing.T) {
	type args struct {
		connectingString string
	}
	tests := []struct {
		name    string
		args    args
		want    FileInfoBase
		wantErr bool
	}{
		{
			"simple",
			args{"File=./file_ib;Locale=ru_RU"},
			FileInfoBase{
				File:   "./file_ib",
				Locale: "ru_RU",
			},
			false,
		},
		{
			"error",
			args{"File ./file_ib2"},
			FileInfoBase{},
			true,
		},
		{
			"full",
			args{"File=./file_ib;Locale=ru_RU;Usr=User;Pwd=Password;Prmod=1;LicDstr=Y"},
			FileInfoBase{
				File:   "./file_ib",
				Locale: "ru_RU",
				InfoBase: InfoBase{
					Usr:     "User",
					Pwd:     "Password",
					LicDstr: true,
					Prmod:   true,
				},
			},
			false,
		},

		{
			"with IBConnectionString",
			args{"/IBConnectionString File=./file_ib"},
			FileInfoBase{
				File: "./file_ib",
			},
			false,
		},
		{
			"with error IBConnectionString",
			args{"/IBConnectionString=File=./file_ib"},
			FileInfoBase{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileInfobaseFromConnectingString(tt.args.connectingString)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileInfobaseFromConnectingString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileInfobaseFromConnectingString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfobaseFromConnectingString(t *testing.T) {
	type args struct {
		connectingString string
	}
	tests := []struct {
		name    string
		args    args
		want    Infobase
		wantErr bool
	}{
		{
			"file",
			args{"File=./file_ib;Locale=ru_RU"},
			FileInfoBase{
				File:   "./file_ib",
				Locale: "ru_RU",
			},
			false,
		},
		{
			"server",
			args{"Srvr=test_server;Ref=test_ib;"},
			ServerInfoBase{
				Srvr: "test_server",
				Ref:  "test_ib",
			},
			false,
		},

		{
			"no file or server",
			args{"FFF=./file_ib;Locale=ru_RU"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InfobaseFromConnectingString(tt.args.connectingString)
			if (err != nil) != tt.wantErr {
				t.Errorf("InfobaseFromConnectingString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InfobaseFromConnectingString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfobaseFromConnectingStringOrNil(t *testing.T) {
	type args struct {
		connectingString string
	}
	tests := []struct {
		name string
		args args
		want Infobase
	}{
		{
			"file",
			args{"File=./file_ib;Locale=ru_RU"},
			FileInfoBase{
				File:   "./file_ib",
				Locale: "ru_RU",
			},
		},
		{
			"server",
			args{"Srvr=test_server;Ref=test_ib;"},
			ServerInfoBase{
				Srvr: "test_server",
				Ref:  "test_ib",
			},
		},

		{
			"no file or server",
			args{"FFF=./file_ib;Locale=ru_RU"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InfobaseFromConnectingStringOrNil(tt.args.connectingString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InfobaseFromConnectingStringOrNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerInfobaseFromConnectingString(t *testing.T) {
	type args struct {
		connectingString string
	}
	tests := []struct {
		name    string
		args    args
		want    ServerInfoBase
		wantErr bool
	}{
		{
			"simple",
			args{"Srvr=test_server;Ref=test_ib;"},
			ServerInfoBase{
				Srvr: "test_server",
				Ref:  "test_ib",
			},
			false,
		},
		{
			"error",
			args{"Srvr=test_server"},
			ServerInfoBase{},
			true,
		},
		{
			"full",
			args{"Srvr=test_server;Ref=test_ib;Usr=User;Pwd=Password;Prmod=1;LicDstr=Y"},
			ServerInfoBase{
				Srvr: "test_server",
				Ref:  "test_ib",
				InfoBase: InfoBase{
					Usr:     "User",
					Pwd:     "Password",
					LicDstr: true,
					Prmod:   true,
				},
			},
			false,
		},

		{
			"with IBConnectionString",
			args{"/IBConnectionString Srvr=test_server;Ref=test_ib"},
			ServerInfoBase{
				Srvr: "test_server",
				Ref:  "test_ib",
			},
			false,
		},
		{
			"with error IBConnectionString",
			args{"/IBConnectionString=Srvr=test_server;Ref=test_ib"},
			ServerInfoBase{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ServerInfobaseFromConnectingString(tt.args.connectingString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServerInfobaseFromConnectingString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerInfobaseFromConnectingString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseDatabaseSeparatorList(t *testing.T) {
	type args struct {
		stringValue string
	}
	tests := []struct {
		name    string
		args    args
		want    DatabaseSeparatorList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDatabaseSeparatorList(tt.args.stringValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDatabaseSeparatorList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDatabaseSeparatorList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
