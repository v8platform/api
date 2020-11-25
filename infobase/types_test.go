package infobase

import (
	"github.com/stretchr/testify/suite"

	"testing"
)

type TypesTestSuite struct {
	baseTestSuite
}

func (b *TypesTestSuite) SetupSuite() {

}

func Test_TypesTestSuite(t *testing.T) {
	suite.Run(t, new(TypesTestSuite))
}

func (t *TypesTestSuite) TestConnectionStringServerBase() {

	tests := []struct {
		name   string
		fields ServerInfoBase
		want   string
	}{
		{
			"simple",
			ServerInfoBase{
				Srvr: "test_server",
				Ref:  "test_base",
			},
			"/IBConnectionString Srvr=test_server;Ref='test_base'",
		},
		{
			"with auth",
			ServerInfoBase{
				Srvr: "test_server",
				Ref:  "test_base",
			}.WithAuth("user", "pwd"),
			"/IBConnectionString Usr=user;Pwd=pwd;Srvr=test_server;Ref='test_base'",
		},
		{
			"with LicDstr",
			ServerInfoBase{
				InfoBase: InfoBase{
					LicDstr: true,
				},
				Srvr: "test_server",
				Ref:  "test_base",
			}.WithAuth("user", "pwd"),
			"/IBConnectionString Usr=user;Pwd=pwd;LicDstr=Y;Srvr=test_server;Ref='test_base'",
		},
		{
			"with Prmod",
			ServerInfoBase{
				InfoBase: InfoBase{
					Prmod: true,
				},
				Srvr: "test_server",
				Ref:  "test_base",
			},
			"/IBConnectionString Prmod=1;Srvr=test_server;Ref='test_base'",
		},
		{
			"with Zn",
			ServerInfoBase{
				InfoBase: InfoBase{
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
			ib := ServerInfoBase{
				InfoBase: tt.fields.InfoBase,
				Srvr:     tt.fields.Srvr,
				Ref:      tt.fields.Ref,
			}

			got := ib.ConnectionString()

			t.r().Equal(got, tt.want, "ConnectionString() = %v, want %v", got, tt.want)

		})
	}

}

func (t *TypesTestSuite) TestConnectionStringFileBase() {

	tests := []struct {
		name   string
		fields FileInfoBase
		want   string
	}{
		{
			"simple",
			FileInfoBase{
				File: "./file_base_dir",
			},
			"/IBConnectionString File='./file_base_dir'",
		},
		{
			"with auth",
			FileInfoBase{
				File: "./file_base_dir",
			}.WithAuth("user", "pwd"),
			"/IBConnectionString Usr=user;Pwd=pwd;File='./file_base_dir'",
		},
		{
			"with LicDstr",
			FileInfoBase{
				InfoBase: InfoBase{
					LicDstr: true,
				},
				File: "./file_base_dir",
			}.WithAuth("user", "pwd"),
			"/IBConnectionString Usr=user;Pwd=pwd;LicDstr=Y;File='./file_base_dir'",
		},
		{
			"with Prmod",
			FileInfoBase{
				InfoBase: InfoBase{
					Prmod: true,
				},
				File: "./file_base_dir",
			},
			"/IBConnectionString Prmod=1;File='./file_base_dir'",
		},
		{
			"with Zn",
			FileInfoBase{
				InfoBase: InfoBase{
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
