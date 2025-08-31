package env

import (
	"fmt"
	"testing"
)

// Test using a command line like:
// MY_STRING_VAR="hello shell" MY_BOOL_VAR=1 go test -v

func TestEnvSet_EnvSetParse(t *testing.T) {

	var myStringVar string
	var myBoolVar bool
	var myIntVar int
	var myFloatVar float64

	envset := EnvSet{}
	envset.StringVar(&myStringVar, "MY_STRING_VAR", "toto", "string test variable")
	envset.BoolVar(&myBoolVar, "MY_BOOL_VAR", false, "boolean test variable")
	envset.IntVar(&myIntVar, "MY_INT_VAR", 1234, "port number")
	envset.FloatVar(&myFloatVar, "MY_FLOAT_VAR", 3.4, "sampling rate")
	envset.Help()

	err := envset.Parse()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(envset)
}

func TestEnvSet_DefaultParse(t *testing.T) {

	var myStringVar string
	var myBoolVar bool
	var myIntVar int
	var myFloatVar float64

	StringVar(&myStringVar, "MY_STRING_VAR", "toto", "string test variable")
	BoolVar(&myBoolVar, "MY_BOOL_VAR", false, "boolean test variable")
	IntVar(&myIntVar, "MY_INT_VAR", 1234, "port number")
	FloatVar(&myFloatVar, "MY_FLOAT_VAR", 3.4, "sampling rate")
	Help()

	err := Parse()
	if err != nil {
		t.Error(err)
	}
	Print()
}
