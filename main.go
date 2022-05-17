package ayika

import (
	"encoding/json"
	"os"
	"strconv"
)

var Load = &load{}

// Get the value of an environment variable.
func Value(key string) string {
	return os.Getenv(key)
}

// Get the value of an environment variable as an integer.
func Num(key string) (int64, error) {
	num, err := strconv.ParseInt(Value(key), 10, 64)
	reportError(err, "Could not parse numerical env variable.")
	return num, err
}

// Get the value of an environment variable as a boolean.
func Bool(key string) bool {
	val := Value(key)
	num, err := strconv.ParseInt(Value(key), 10, 64)
	if err == nil {
		return num != 0
	}

	return len(val) > 0 && val != "" && val != "false"
}

// Get the value of an environment variable as a float.
func Float(key string) (float64, error) {
	num, err := strconv.ParseFloat(Value(key), 64)
	reportError(err, "Could not parse float env variable.")
	return num, err
}

// Get the value of an environment variable as a JSON interface.
func JSON[V interface{}](key string) (V, error) {
	var v V
	err := json.Unmarshal([]byte(Value(key)), &v)
	reportError(err, "Could not parse JSON env variable.")
	return v, err
}
