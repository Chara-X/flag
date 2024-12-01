package flag

import (
	"fmt"
	"strings"
)

func ExampleFlagSet() {
	var osArgs = strings.Fields("root -g global sub -l local arg")
	var global string
	var root = NewFlagSet(osArgs[0])
	root.Var(newStringValue(&global), "g", "This is a global flag")
	root.Parse(osArgs[1:])
	var args = root.Args()
	fmt.Println("root:", root.Name())
	fmt.Println("global:", global)
	switch args[0] {
	case "sub":
		var local string
		var sub = NewFlagSet(args[0])
		sub.Var(newStringValue(&local), "l", "This is local flag")
		sub.Parse(args[1:])
		fmt.Println("sub:", sub.Name())
		fmt.Println("local:", local)
		fmt.Println("args:", sub.Args())
	default:
		fmt.Println("args:", args)
	}
	// Output:
	// root: root
	// global: global
	// sub: sub
	// local: local
	// args: [arg]
}

type stringValue string

func newStringValue(p *string) *stringValue { return (*stringValue)(p) }
func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}
func (s *stringValue) String() string { return string(*s) }
