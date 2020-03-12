package v8

import "github.com/Khorevaa/go-v8runner/designer"

func DumpIB(file string) designer.DumpIBOptions {

	command := designer.DumpIBOptions{
		Designer: designer.NewDesigner(),
		File:     file,
	}

	return command
}

func RestoreIB(file string) designer.RestoreIBOptions {

	command := designer.RestoreIBOptions{
		Designer: designer.NewDesigner(),
		File:     file,
	}

	return command
}

func IBRestoreIntegrity() designer.IBRestoreIntegrityOptions {

	return designer.IBRestoreIntegrityOptions{
		Designer: designer.NewDesigner(),
	}
}
