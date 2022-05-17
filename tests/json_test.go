package ayika_test

import (
	"testing"

	"github.com/ededejr/ayika"
)

func TestJSON(t *testing.T) {
	t.Setenv("integer", "5")
	t.Setenv("string", "golang")
	t.Setenv("bool", "true")
	t.Setenv("float", "3.14")
	t.Setenv("json", "{\"key\":\"value\"}")

	type KeyValue struct {
		Key string `json:"key"`
	}

	cases := []testCase[string, KeyValue]{
		{"integer", "Integers parsed should throw an error", func(t *testing.T, result KeyValue, err error) bool {
			if err == nil {
				t.Error("Expected error, got none")
				return false
			}
			return true
		}},
		{"float", "Floats parsed should throw an error", func(t *testing.T, result KeyValue, err error) bool {
			if err == nil {
				t.Error("Expected error, got none")
				return false
			}
			return true
		}},
		{"string", "Strings parsed should throw an error", func(t *testing.T, result KeyValue, err error) bool {
			if err == nil {
				t.Error("Expected error, got none")
				return false
			}
			return true
		}},
		{"bool", "Booleans parsed should throw an error", func(t *testing.T, result KeyValue, err error) bool {
			if err == nil {
				t.Error("Expected error, got none")
				return false
			}
			return true
		}},
		{"json", "JSON parsed should return value", func(t *testing.T, result KeyValue, err error) bool {
			if result.Key != "value" || err != nil {
				t.Errorf("Expected Key to be \"value\", got %s", result)
				return false
			}
			return true
		}},
	}

	runTestCasesWithErrors(t, "JSON()", ayika.JSON[KeyValue], cases)
}
