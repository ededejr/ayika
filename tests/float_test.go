package ayika_test

import (
	"testing"

	"github.com/ededejr/ayika"
)

func TestFloat(t *testing.T) {
	t.Setenv("integer", "5")
	t.Setenv("string", "golang")
	t.Setenv("bool", "true")
	t.Setenv("float", "3.14")

	cases := []testCase[string, float64]{
		{"integer", "Integers are parsed as float64", func(t *testing.T, result float64, err error) bool {
			expectedValue := float64(5)
			if result != expectedValue {
				t.Errorf("Expected %f, got %f", expectedValue, result)
				return false
			}
			return true
		}},
		{"float", "Floats are parsed as float64", func(t *testing.T, result float64, err error) bool {
			expectedValue := float64(3.14)
			if result != expectedValue {
				t.Errorf("Expected %f, got %f", expectedValue, result)
				return false
			}
			return true
		}},
		{"string", "Strings parsed should throw an error", func(t *testing.T, result float64, err error) bool {
			if err == nil {
				t.Error("Expected error, got none")
				return false
			}
			return true
		}},
		{"bool", "Booleans parsed should throw an error", func(t *testing.T, result float64, err error) bool {
			if err == nil {
				t.Error("Expected error, got none")
				return false
			}
			return true
		}},
	}

	runTestCasesWithErrors(t, "Float()", ayika.Float, cases)
}
