package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type envtype interface{ int | bool | string | float64 }

// envvar holds the definition of an environment variable. It is a
// generic (template) so that it can be used for variables of different
// types
type envvar[T envtype] struct {
	value    *T
	name     string
	vdefault T
	help     string
}

type EnvSet struct {
	boolvars   []envvar[bool]
	stringvars []envvar[string]
	intvars    []envvar[int]
	floatvars  []envvar[float64]
}

func (s *EnvSet) IntVar(v *int, vname string, vdefault int, vhelp string) {
	s.intvars = append(s.intvars, envvar[int]{
		value:    v,
		name:     vname,
		vdefault: vdefault,
		help:     vhelp,
	})
}

func (s *EnvSet) BoolVar(v *bool, vname string, vdefault bool, vhelp string) {
	s.boolvars = append(s.boolvars, envvar[bool]{
		value:    v,
		name:     vname,
		vdefault: vdefault,
		help:     vhelp,
	})
}

func (s *EnvSet) StringVar(v *string, vname string, vdefault string, vhelp string) {
	s.stringvars = append(s.stringvars, envvar[string]{
		value:    v,
		name:     vname,
		vdefault: vdefault,
		help:     vhelp,
	})
}

func (s *EnvSet) FloatVar(v *float64, vname string, vdefault float64, vhelp string) {
	s.floatvars = append(s.floatvars, envvar[float64]{
		value:    v,
		name:     vname,
		vdefault: vdefault,
		help:     vhelp,
	})
}

// ---------------------------------------------------------------------------
// Functions to parse the values of environment variables

var _string2bool_truelist = []string{"y", "o", "yes", "oui", "vrai", "true", "1"}

func string2bool(v string) (bool, error) {
	lowvalue := strings.ToLower(v)
	for _, v := range _string2bool_truelist {
		if lowvalue == v {
			return true, nil
		}
	}
	return false, nil
}

func string2float(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

func parseList[T envtype](t []envvar[T], parseVar func(string) (T, error)) error {
	for _, v := range t {
		value, exists := os.LookupEnv(v.name)
		if exists {
			r, err := parseVar(value)
			if err != nil {
				*v.value = v.vdefault
				return fmt.Errorf("the value for variable %s can not be parsed", v.name)
			}
			*v.value = r
		} else {
			*v.value = v.vdefault
		}
	}
	return nil
}

// ---------------------------------------------------------------------------
// Global parsing process (process that read the variable values from
// the environment)

func (s EnvSet) Parse() error {
	err := parseList(s.stringvars, func(s string) (string, error) { return s, nil })
	if err != nil {
		return err
	}
	err = parseList(s.boolvars, string2bool)
	if err != nil {
		return err
	}
	err = parseList(s.intvars, strconv.Atoi)
	if err != nil {
		return err
	}
	err = parseList(s.floatvars, string2float)
	if err != nil {
		return err
	}

	return nil
}

func (s EnvSet) Help() {
	for _, v := range s.stringvars {
		fmt.Printf("%-16s: %s (%T, %s)\n", v.name, v.help, v.vdefault, v.vdefault)
	}
	for _, v := range s.boolvars {
		fmt.Printf("%-16s: %s (%T, %t)\n", v.name, v.help, v.vdefault, v.vdefault)
	}
	for _, v := range s.intvars {
		fmt.Printf("%-16s: %s (%T, %d)\n", v.name, v.help, v.vdefault, v.vdefault)
	}
	for _, v := range s.floatvars {
		fmt.Printf("%-16s: %s (%T, %.4f)\n", v.name, v.help, v.vdefault, v.vdefault)
	}
}

func (s EnvSet) String() string {
	str := ""
	for _, v := range s.stringvars {
		str += fmt.Sprintf("%-16s: %s\n", v.name, *v.value)
	}
	for _, v := range s.boolvars {
		str += fmt.Sprintf("%-16s: %t\n", v.name, *v.value)
	}
	for _, v := range s.intvars {
		str += fmt.Sprintf("%-16s: %d\n", v.name, *v.value)
	}
	for _, v := range s.floatvars {
		str += fmt.Sprintf("%-16s: %.4f\n", v.name, *v.value)
	}
	return str
}

// ---------------------------------------------------------------------------
// Default environment variable set

var defaultEnvSet EnvSet

func BoolVar(v *bool, vname string, vdefault bool, vhelp string) {
	defaultEnvSet.BoolVar(v, vname, vdefault, vhelp)
}

func StringVar(v *string, vname string, vdefault string, vhelp string) {
	defaultEnvSet.StringVar(v, vname, vdefault, vhelp)
}

func IntVar(v *int, vname string, vdefault int, vhelp string) {
	defaultEnvSet.IntVar(v, vname, vdefault, vhelp)
}

func FloatVar(v *float64, vname string, vdefault float64, vhelp string) {
	defaultEnvSet.FloatVar(v, vname, vdefault, vhelp)
}

func Help() {
	defaultEnvSet.Help()
}

func Print() {
	fmt.Println(defaultEnvSet)
}

func Parse() error {
	return defaultEnvSet.Parse()
}
