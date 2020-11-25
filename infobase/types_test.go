package infobase

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"reflect"

	"testing"
)

type TypesTestSuite struct {
	suite.Suite
}

func (b *TypesTestSuite) SetupSuite() {

}

func (s *TypesTestSuite) r() *require.Assertions {
	return s.Require()
}

func Test_TypesTestSuite(t *testing.T) {
	suite.Run(t, new(TypesTestSuite))
}

func (t *TypesTestSuite) TestConnectionStringServerBase() {

	tests := []struct {
		name   string
		fields Server
		want   string
	}{
		{
			"simple",
			Server{
				Srvr: "test_server",
				Ref:  "test_base",
			},
			"/IBConnectionString Srvr=test_server;Ref='test_base'",
		},
		{
			"with auth",
			Server{
				Common: Common{
					Usr: "user",
					Pwd: "pwd",
				},
				Srvr: "test_server",
				Ref:  "test_base",
			},
			"/IBConnectionString Usr=user;Pwd=pwd;Srvr=test_server;Ref='test_base'",
		},
		{
			"with LicDstr",
			Server{
				Common: Common{
					LicDstr: true,
					Usr:     "user",
					Pwd:     "pwd",
				},
				Srvr: "test_server",
				Ref:  "test_base",
			},
			"/IBConnectionString Usr=user;Pwd=pwd;LicDstr=Y;Srvr=test_server;Ref='test_base'",
		},
		{
			"with Prmod",
			Server{
				Common: Common{
					Prmod: true,
				},
				Srvr: "test_server",
				Ref:  "test_base",
			},
			"/IBConnectionString Prmod=1;Srvr=test_server;Ref='test_base'",
		},
		{
			"with Zn",
			Server{
				Common: Common{
					Zn: DatabaseSeparatorList{
						DatabaseSeparator{
							Use:   true,
							Value: "first",
						},
						DatabaseSeparator{
							Use:   false,
							Value: "second",
						},
					},
				},
				Srvr: "test_server",
				Ref:  "test_base",
			},
			"/IBConnectionString ZN=+first,-second;Srvr=test_server;Ref='test_base'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func() {
			ib := Server{
				Common: tt.fields.Common,
				Srvr:   tt.fields.Srvr,
				Ref:    tt.fields.Ref,
			}

			got := ib.ConnectionString()

			t.r().Equal(got, tt.want, "ConnectionString() = %v, want %v", got, tt.want)

		})
	}

}

func (t *TypesTestSuite) TestConnectionStringFileBase() {

	tests := []struct {
		name   string
		fields File
		want   string
	}{
		{
			"simple",
			File{
				File: "./file_base_dir",
			},
			"/IBConnectionString File='./file_base_dir'",
		},
		{
			"with auth",
			File{
				File: "./file_base_dir",
				Common: Common{
					Usr: "user",
					Pwd: "pwd",
				},
			},
			"/IBConnectionString Usr=user;Pwd=pwd;File='./file_base_dir'",
		},
		{
			"with LicDstr",
			File{
				Common: Common{
					Usr:     "user",
					Pwd:     "pwd",
					LicDstr: true,
				},
				File: "./file_base_dir",
			},
			"/IBConnectionString Usr=user;Pwd=pwd;LicDstr=Y;File='./file_base_dir'",
		},
		{
			"with Prmod",
			File{
				Common: Common{
					Prmod: true,
				},
				File: "./file_base_dir",
			},
			"/IBConnectionString Prmod=1;File='./file_base_dir'",
		},
		{
			"with Zn",
			File{
				Common: Common{
					Zn: DatabaseSeparatorList{
						DatabaseSeparator{
							Use:   true,
							Value: "first",
						},
						DatabaseSeparator{
							Use:   false,
							Value: "second",
						},
					},
				},
				File: "./file_base_dir",
			},
			"/IBConnectionString ZN=+first,-second;File='./file_base_dir'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func() {

			got := tt.fields.ConnectionString()

			t.r().Equal(got, tt.want, "ConnectionString() = %v, want %v", got, tt.want)

		})
	}

}

func TestWithAuth(t *testing.T) {
	type args struct {
		ib       Infobase
		user     string
		password string
	}
	tests := []struct {
		name string
		args args
		want Infobase
	}{
		{
			"simple file",
			args{
				File{},
				"user",
				"pwd",
			},
			File{
				Common: Common{
					Usr: "user",
					Pwd: "pwd",
				},
			},
		},
		{
			"simple server",
			args{
				Server{},
				"user",
				"pwd",
			},
			Server{
				Common: Common{
					Usr: "user",
					Pwd: "pwd",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithAuth(tt.args.ib, tt.args.user, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
