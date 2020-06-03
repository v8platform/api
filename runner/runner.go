package runner

import (
	"fmt"
	"github.com/khorevaa/go-v8platform/find"
	"github.com/khorevaa/go-v8platform/runner/cmd"
	"github.com/khorevaa/go-v8platform/types"
	"strings"

	"context"
	"github.com/khorevaa/go-v8platform/errors"
)

var defaultVersion = "8.3"

type v8Runner struct {
	Options   *Options
	Where     types.InfoBase
	What      types.Command
	ctx       context.Context
	commandV8 string
}

func newRunner(ctx context.Context, where types.InfoBase, what types.Command, opts ...interface{}) v8Runner {

	options := defaultOptions()

	inlineOptions := getOptions(opts...)
	if inlineOptions != nil {
		options = inlineOptions
	}

	o := clearOpts(opts)

	options.Options(o...)

	r := v8Runner{
		Where:   where,
		What:    what,
		Options: options,
		ctx:     ctx,
	}

	return r
}

func Run(where types.InfoBase, what types.Command, opts ...interface{}) error {

	ctx := context.Background()

	p, err := Background(ctx, where, what, opts...)

	if err != nil {
		return err
	}

	return <-p.Wait()
}

func Background(ctx context.Context, where types.InfoBase, what types.Command, opts ...interface{}) (Process, error) {

	r := newRunner(ctx, where, what, opts...)

	err := checkCommand(r.What)

	if err != nil {
		return nil, err
	}

	r.commandV8, err = getV8Path(*r.Options)

	if err != nil {
		return nil, err
	}

	p := r.run()

	return p, nil

}

func (r *v8Runner) run() Process {

	args := getCmdArgs(r.Where, r.What, *r.Options)

	runner := prepareRunner(r.commandV8, args, *r.Options)

	p := background(runner, r.ctx)

	return p

}

func checkCommand(what types.Command) (err error) {
	err = what.Check()
	return
}

func getCmdArgs(where types.InfoBase, what types.Command, options Options) []string {

	var args []string

	values := types.NewValues()

	connectString := where.Values()

	if what.Command() == types.COMMAND_CREATEINFOBASE {
		connectString.Append(*what.Values())
	} else {
		values.Set("/IBConnectionString", types.SpaceSep,
			fmt.Sprintf("%s;", strings.Join(connectString.Values(), ";")))
		values.Append(*what.Values())
	}

	values.Append(options.commonValues)
	values.Append(*options.Values())
	values.Append(options.customValues)

	args = append(args, what.Command())
	if what.Command() == types.COMMAND_CREATEINFOBASE {

		args = append(args, strings.Join(connectString.Values(), ";"))

		clearValuesForCreateInfobase(values)
	}
	args = append(args, values.Values()...)

	return args
}

func prepareRunner(command string, args []string, options Options) Runner {

	r := cmd.NewCmdRunner(command, args,
		cmd.WithContext(options.Context),
		cmd.WithOutFilePath(options.Out),
		cmd.WithDumpResultFilePath(options.DumpResult),
	)

	return r
}

func getV8Path(options Options) (string, error) {
	if len(options.v8path) > 0 {
		return options.v8path, nil
	}

	v8 := defaultVersion
	if len(options.Version) > 0 {
		v8 = options.Version
	}

	v8path, err := find.PlatformByVersion(v8, find.WithBitness(find.V8_x64x32))

	if err != nil {

		err = errors.NotExist.Newf("Version %s not found", options.Version)
		_ = errors.AddErrorContext(err, "version", options.Version)

		return "", err
	}

	return v8path, nil

}

func clearValuesForCreateInfobase(v *types.Values) {

	// TODO Сделать очистку значение

}

func defaultOptions() *Options {

	options := Options{}

	options.NewOutFile()
	options.NewDumpResultFile()
	//options.customValues = *types.NewValues()
	//options.commonValues = *types.NewValues()

	return &options
}

func getOptions(opts ...interface{}) *Options {

	for _, opt := range opts {

		switch opt.(type) {

		case Options:
			userOptions, _ := opt.(Options)
			return &userOptions
		case *Options:
			userOptions, _ := opt.(*Options)
			return userOptions
		}

	}

	return nil
}

func clearOpts(opts []interface{}) []Option {

	var o []Option

	for _, opt := range opts {

		if fn, ok := opt.(Option); ok {
			o = append(o, fn)
		}
	}
	return o
}
