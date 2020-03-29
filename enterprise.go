package v8

import "github.com/Khorevaa/go-v8runner/enterprise"

func Execute(file string) enterprise.ExecuteOptions {

	command := enterprise.ExecuteOptions{
		Enterprise: enterprise.NewEnterprise(),
		File:       file,
	}

	return command
}
