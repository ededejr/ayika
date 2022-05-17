package ayika_test

import (
	"testing"

	"github.com/ededejr/ayika"
)

func TestValue(t *testing.T) {
	t.Setenv("num", "1")
	t.Setenv("string", "golang")
	t.Setenv("bool", "true")
	t.Setenv("json", "{\"key\":\"value\"}")

	cases := []testCase[string, string]{
		{"num", "Numeric values are parsed as strings", func(t *testing.T, result string, err error) bool {
			if result != "1" {
				t.Errorf("Expected %s, got %s", "1", result)
				return false
			}
			return true
		}},
		{"string", "String values are parsed as strings", func(t *testing.T, result string, err error) bool {
			if result != "golang" {
				t.Errorf("Expected %s, got %s", "string", result)
				return false
			}
			return true
		}},
		{"bool", "Boolean values are parsed as strings", func(t *testing.T, result string, err error) bool {
			if result != "true" {
				t.Errorf("Expected %s, got %s", "true", result)
				return false
			}
			return true
		}},
		{"json", "JSON values are parsed as strings", func(t *testing.T, result string, err error) bool {
			if result != "{\"key\":\"value\"}" {
				t.Errorf("Expected %s, got %s", "{\"key\":\"value\"}", result)
				return false
			}
			return true
		}},
	}

	runTestCases(t, "Value()", ayika.Value, cases)
}
