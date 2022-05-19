package ayika_test

import (
	"testing"

	"github.com/ededejr/ayika"
)

func TestNum(t *testing.T) {
	t.Setenv("digit", "1")
	t.Setenv("string", "golang")
	t.Setenv("bool", "true")
	t.Setenv("multidigit", "100")

	cases := []testCase[string, int64]{
		{"digit", "Single digits are parsed as int64", func(t *testing.T, result int64, err error) bool {
			expectedValue := int64(1)
			if result != expectedValue {
				t.Errorf("Expected %d, got %d", expectedValue, result)
				return false
			}
			return true
		}},
		{"multidigit", "Multiple digits are parsed as int64", func(t *testing.T, result int64, err error) bool {
			expectedValue := int64(100)
			if result != expectedValue {
				t.Errorf("Expected %d, got %d", expectedValue, result)
				return false
			}
			return true
		}},
		{"string", "Strings parsed should throw an error", func(t *testing.T, result int64, err error) bool {
			if err == nil {
				t.Error("Expected error, got none")
				return false
			}
			return true
		}},
		{"bool", "Booleans parsed should throw an error", func(t *testing.T, result int64, err error) bool {
			if err == nil {
				t.Error("Expected error, got none")
				return false
			}
			return true
		}},
	}

	runTestCasesWithErrors(t, "Num()", ayika.Num, cases)
}
