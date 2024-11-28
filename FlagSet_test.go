package flag

import "fmt"

func ExampleFlagSet() {
	var f = NewFlagSet("curl")
	var method = "GET"
	f.Var((*stringValue)(&method), "X", "HTTP method")
	f.Parse([]string{"-X", "POST", "http://example.com"})
	fmt.Println(method)
	// Output:
	// POST
}

type stringValue string

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}
func (s *stringValue) String() string { return string(*s) }
