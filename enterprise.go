package v8

import "github.com/v8platform/enterprise"

func Execute(file string) enterprise.ExecuteOptions {

	command := enterprise.ExecuteOptions{
		File: file,
	}

	return command
}
