package flag

import "flag"

type FlagSet struct {
	f     flag.FlagSet
	name  string
	flags map[string]*flag.Flag
	args  []string
	Usage func()
}

func NewFlagSet(name string) *FlagSet {
	if Reference {
		return &FlagSet{f: *flag.NewFlagSet(name, flag.PanicOnError)}
	}
	return &FlagSet{name: name, flags: map[string]*flag.Flag{}, Usage: func() {}}
}
func (f *FlagSet) Var(value flag.Value, name string, usage string) {
	if Reference {
		f.f.Var(value, name, usage)
		return
	}
	f.flags[name] = &flag.Flag{Value: value, Name: name, Usage: usage, DefValue: value.String()}
}
func (f *FlagSet) Parse(arguments []string) error {
	if Reference {
		return f.f.Parse(arguments)
	}
	f.args = arguments
	for {
		if len(f.args) == 0 {
			break
		}
		var arg = f.args[0]
		if len(arg) < 2 || arg[0] != '-' {
			break
		}
		f.args = f.args[1:]
		if arg == "--" {
			break
		}
		var hyphen = 1
		if arg[1] == '-' {
			hyphen = 2
		}
		var name = arg[hyphen:]
		if name == "help" {
			f.Usage()
			break
		}
		switch v := f.flags[name].Value.(type) {
		case interface {
			flag.Value
			IsBoolFlag() bool
		}:
			v.Set("true")
		default:
			v.Set(f.args[0])
			f.args = f.args[1:]
		}
	}
	return nil
}
func (f *FlagSet) Args() []string {
	if Reference {
		return f.Args()
	}
	return f.args
}
