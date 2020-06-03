package v8

import "github.com/khorevaa/go-v8platform/enterprise"

func Execute(file string) enterprise.ExecuteOptions {

	command := enterprise.ExecuteOptions{
		Enterprise: enterprise.NewEnterprise(),
		File:       file,
	}

	return command
}
