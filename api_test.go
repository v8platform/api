package v8

import (
	"github.com/stretchr/testify/assert"
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

func TestApiConfig(t *testing.T) {

	ib := NewFileInfobase("./test_ib")

	tests := []struct {
		name     string
		where    ConnectionString
		want     Command
		opts     []interface{}
		wantArgs []string
	}{
		{
			"LoadCfg",
			ib,
			LoadCfg("./1cv8.cf"),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/LoadCfg ./1cv8.cf",
			},
		},
		{
			"LoadCfg_UpdateDBCfg",
			ib,
			LoadCfg("./1cv8.cf", UpdateDBCfg(true, true)),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/LoadCfg ./1cv8.cf",
				"/UpdateDBCfg",
				//"-Dynamic+", FIX: поправить после исправления https://github.com/v8platform/marshaler/issues/1
				"-Server",
			},
		},
		{
			"LoadConfigFromFiles",
			ib,
			LoadConfigFromFiles("./src"),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/LoadConfigFromFiles ./src",
			},
		},
		{
			"LoadConfigFromFiles_WithFiles",
			ib,
			LoadConfigFromFiles("./src").WithFiles("./src/file1.xml", "./src/file2.xml"),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/LoadConfigFromFiles ./src",
				"-files ./src/file1.xml,./src/file2.xml",
			},
		},
		{
			"LoadConfigFromFiles_WithListFile",
			ib,
			LoadConfigFromFiles("./src").WithListFile("./file_list.xml"),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/LoadConfigFromFiles ./src",
				"-listFile ./file_list.xml",
			},
		},
		{
			"LoadConfigFromFiles_WithUpdateDumpInfo",
			ib,
			LoadConfigFromFiles("./src").WithUpdateDumpInfo(),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/LoadConfigFromFiles ./src",
				"-updateConfigDumpInfo",
			},
		},
		{
			"LoadConfigFromFiles_UpdateDBCfg",
			ib,
			LoadConfigFromFiles("./src", UpdateDBCfg(true, true)),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/LoadConfigFromFiles ./src",
				"/UpdateDBCfg",
				//"-Dynamic+", FIX: поправить после исправления https://github.com/v8platform/marshaler/issues/1
				"-Server",
			},
		},
		{
			"UpdateCfg",
			ib,
			UpdateCfg("./1cv8.cf", false),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/UpdateCfg ./1cv8.cf",
			},
		},
		{
			"UpdateCfg_force",
			ib,
			UpdateCfg("./1cv8.cf", true),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/UpdateCfg ./1cv8.cf",
				"-Force",
			},
		},
		{
			"UpdateCfg_UpdateDBCfg",
			ib,
			UpdateCfg("./1cv8.cf", true, UpdateDBCfg(true, true)),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/UpdateCfg ./1cv8.cf",
				"-Force",
				"/UpdateDBCfg",
				//"-Dynamic+", FIX: поправить после исправления https://github.com/v8platform/marshaler/issues/1
				"-Server",
			},
		},
		{
			"DumpCfg",
			ib,
			DumpCfg("./1cv8.cf"),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/DumpCfg ./1cv8.cf",
			},
		},
		{
			"DumpConfigToFiles",
			ib,
			DumpConfigToFiles("./src", false),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/DumpConfigToFiles ./src",
			},
		},
		{
			"DumpConfigToFiles_force",
			ib,
			DumpConfigToFiles("./src", true),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/DumpConfigToFiles ./src",
				"-force",
			},
		},
		{
			"DumpConfigToFiles_update",
			ib,
			DumpConfigToFiles("./src", false).WithUpdate(""),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/DumpConfigToFiles ./src",
				"-update",
			},
		},
		{
			"DumpConfigToFiles_configDumpInfo",
			ib,
			DumpConfigToFiles("./src", false).WithUpdate("./dumpInfo.xml"),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/DumpConfigToFiles ./src",
				"-update",
				"-configDumpInfoForChanges ./dumpInfo.xml",
			},
		},
		{
			"GetChangesForConfigDump",
			ib,
			GetChangesForConfigDump("./src", "./dumpInfo.xml", false),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/DumpConfigToFiles ./src",
				"-getChanges ./dumpInfo.xml",
			},
		},
		{
			"GetChangesForConfigDump_force",
			ib,
			GetChangesForConfigDump("./src", "./dumpInfo.xml", true),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/DumpConfigToFiles ./src",
				"-getChanges ./dumpInfo.xml",
				"-force",
			},
		},
		{
			"GetChangesForConfigDump_WithConfigDumpInfo",
			ib,
			GetChangesForConfigDump("./src", "./dumpInfo.xml", false).
				WithConfigDumpInfo("./old_dumpInfo.xml"),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/DumpConfigToFiles ./src",
				"-getChanges ./dumpInfo.xml",
				"-configDumpInfoForChanges ./old_dumpInfo.xml",
			},
		},
		{
			"DisableCfgSupport",
			ib,
			DisableCfgSupport(),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/ManageCfgSupport",
				"-disableSupport",
			},
		},
		{
			"DisableCfgSupport_force",
			ib,
			DisableCfgSupport(true),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/ManageCfgSupport",
				"-disableSupport",
				"-force",
			},
		},
		{
			"RollbackCfg",
			ib,
			RollbackCfg(),
			nil,
			[]string{
				"DESIGNER",
				"/IBConnectionString File='./test_ib';",
				"/DisableStartupDialogs",
				"/RollbackCfg",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			runTestCase(t, tt)

		})
	}

}

func TestApiOptions(t *testing.T) {

	//ib := NewFileInfobase("./test_ib")

	tests := []struct {
		name     string
		where    ConnectionString
		want     Command
		opts     []interface{}
		wantArgs []string
	}{
		{
			"options",
			Infobase{
				Connect: ServerPath{
					Server: "app-server",
					Ref:    "ib_name",
				},
				User:                "administrator",
				Password:            "password",
				AllowServerLicenses: true,
				SeparatorList: DatabaseSeparatorList{DatabaseSeparator{
					Use:   true,
					Value: "sep1",
				}},
				UsePrivilegedMode: true,
				Locale:            "ru_RU",
			},
			designer.LoadCfgOptions{
				File:     "./1cv8.cf",
				Designer: designer.NewDesigner(),
			},
			[]interface{}{
				WithCredentials("admin", "pwd"),
				WithOut("./out_file", true),
				WithUC("UnlockCode"),
			},
			[]string{
				"DESIGNER",
				"/IBConnectionString Srvr=app-server;Ref='ib_name';Usr=administrator;Pwd=password;LicDstr=Y;ZN=+sep1;Prmod=1;Locale=ru_RU",
				"/DisableStartupDialogs",
				"/LoadCfg ./1cv8.cf",
				"/Out ./out_file -NoTruncate",
				"/UC UnlockCode",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			runTestCase(t, tt)

		})
	}

}

func runTestCase(t *testing.T, tt struct {
	name     string
	where    ConnectionString
	want     Command
	opts     []interface{}
	wantArgs []string
}) {
	v8run := runner.NewPlatformRunner(tt.where, tt.want, tt.opts...)
	got := v8run.Args()

	for _, arg := range tt.wantArgs {
		assert.Contains(t, got, arg)
	}

}
