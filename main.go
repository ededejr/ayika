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
func Num(key string) int64 {
	num, err := strconv.ParseInt(Value(key), 10, 64)
	throwOnError(err, "Could not parse numerical env variable.")
	return num
}

// Get the value of an environment variable as a boolean.
func Bool(key string) bool {
	val := Value(key)
	return len(val) > 0 && val != "" && val != "false"
}

// Get the value of an environment variable as a float.
func Float(key string) float64 {
	num, err := strconv.ParseFloat(Value(key), 64)
	throwOnError(err, "Could not parse numerical float env variable.")
	return num
}

// Get the value of an environment variable as a JSON interface.
func JSON[V interface{}](key string, v V) {
	err := json.Unmarshal([]byte(Value(key)), v)
	throwOnError(err, "Could not parse json env variable.")
}
