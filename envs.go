package envs

import (
	"fmt"
	"os"
	"strconv"
)

type ErrEnvKeyNotFound struct {
	key string
}

func (e ErrEnvKeyNotFound) Error() string {
	return fmt.Sprintf("envs: could not found key for %s", e.key)
}

type ErrEnvValueConvFailed struct {
	key      string
	value    any
	convType string
	must     bool
}

func (e ErrEnvValueConvFailed) Error() string {
	return fmt.Sprintf("envs: cannot convert %s type key for %s, value %s", e.convType, e.key, e.value)
}

func get(key string, defaultValue any) (any, bool) {
	value, found := os.LookupEnv(key)
	if !found || len(value) == 0 {
		return defaultValue, found
	}
	return value, found
}

func mustGet(key string) any {
	value, found := os.LookupEnv(key)
	if !found || len(value) == 0 {
		panic(fmt.Sprintf("%s env value must set!", key))
	}
	return value
}

func MustGetString(key string) string {
	value := mustGet(key)
	resultValue, ok := value.(string)
	if !ok {
		panic(ErrEnvValueConvFailed{key: key, value: value, convType: "string", must: true})
	}
	return resultValue
}

func GetString(key, defaultValue string) (string, error) {
	value, _ := get(key, defaultValue)
	return value.(string), nil
}

func GetInt(key string, defaultValue int) (int, error) {
	value, _ := get(key, defaultValue)

	var (
		resultValue int
		err         error
	)

	switch value := value.(type) {
	case int:
		resultValue = value
	case string:
		if resultValue, err = strconv.Atoi(value); err != nil {
			err = ErrEnvValueConvFailed{key: key, value: value, convType: "int"}
		}
	}
	return resultValue, err
}

func MustGetInt(key string) int {
	value := mustGet(key)
	resultValue, err := strconv.Atoi(value.(string))
	if err != nil {
		panic(ErrEnvValueConvFailed{key: key, value: value, convType: "int", must: true})
	}
	return resultValue
}

func GetBool(key string, defaultValue bool) (bool, error) {
	value, _ := get(key, defaultValue)
	var (
		resultValue bool
		err         error
	)

	switch value := value.(type) {
	case bool:
		resultValue = value
	case string:
		if resultValue, err = strconv.ParseBool(value); err != nil {
			err = ErrEnvValueConvFailed{key: key, value: value, convType: "bool"}
		}
	}
	return resultValue, err
}

func MustGetBool(key string) bool {
	value := mustGet(key)
	resultValue, err := strconv.ParseBool(value.(string))
	if err != nil {
		panic(ErrEnvValueConvFailed{key: key, value: value, convType: "bool", must: true})
	}
	return resultValue
}
