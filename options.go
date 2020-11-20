package v8

import (
	"context"
	"github.com/v8platform/runner"
)

func WithTimeout(timeout int64) runner.Option {
	return runner.WithTimeout(timeout)
}

func WithContext(ctx context.Context) runner.Option {
	return runner.WithContext(ctx)
}

func WithOut(file string, noTruncate bool) runner.Option {
	return runner.WithOut(file, noTruncate)
}

func WithPath(path string) runner.Option {
	return runner.WithPath(path)
}

func WithDumpResult(file string) runner.Option {
	return runner.WithDumpResult(file)
}

func WithVersion(version string) runner.Option {
	return runner.WithVersion(version)
}

func WithCommonValues(cv []string) runner.Option {
	return runner.WithCommonValues(cv)
}

func WithCredentials(user, password string) runner.Option {

	return runner.WithCredentials(user, password)
}

func WithUnlockCode(uc string) runner.Option {

	return runner.WithUnlockCode(uc)
}

func WithUC(uc string) runner.Option {

	return WithUnlockCode(uc)
}
