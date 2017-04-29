package types

import (
	"reflect"
	"strings"

	"github.com/ellcrys/util"
)

// Env represents a collection of environment variables
type Env map[string]string

// NewEnv creates a new Env type from values in map[string]interface{}.
// Non-string values are omitted.
func NewEnv(env map[string]interface{}) Env {
	_env := Env{}
	for k, v := range env {
		if valStr, ok := v.(string); ok {
			_env[k] = valStr
		}
	}
	return _env
}

// Eql checks whether the object is equal to another
func (e Env) Eql(o Env) bool {
	return reflect.DeepEqual(e, o)
}

// ToJSON returns a json encoded representation of the ENV
func (e Env) ToJSON() []byte {
	b, _ := util.ToJSON(e)
	return b
}

// GetFlags returns the flags from an environment variable
func GetFlags(v string) []string {
	var varParts = strings.Split(v, "@")
	var flags = []string{}
	if len(varParts) > 1 && strings.TrimSpace(varParts[1]) != "" {
		flags = strings.Split(varParts[1], ",")
		for i, f := range flags {
			flags[i] = strings.TrimSpace(f)
		}
	}
	return flags
}

// Process applies the flags and returns both public
// environment variables and private variables
func (e Env) Process() (Env, Env) {
	var public = make(Env)
	var private = make(Env)

	for k, v := range e {
		flags := GetFlags(k)
		isPrivate := util.InStringSlice(flags, "private")
		for _, f := range flags {
			switch f {
			case "genRand16":
				v = util.CryptoRandKey(16)
			case "genRand24":
				v = util.CryptoRandKey(24)
			case "genRand32":
				v = util.CryptoRandKey(32)
			case "genRand64":
				v = util.CryptoRandKey(64)
			case "genRand128":
				v = util.CryptoRandKey(128)
			case "genRand256":
				v = util.CryptoRandKey(256)
			case "genRand512":
				v = util.CryptoRandKey(512)
			}
		}
		if isPrivate {
			private[strings.Split(k, "@")[0]] = v
		} else {
			public[strings.Split(k, "@")[0]] = v
		}
	}

	return public, private
}
