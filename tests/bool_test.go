package ayika_test

import (
	"testing"

	"github.com/ededejr/ayika"
)

func TestBool(t *testing.T) {
	t.Setenv("num", "1")
	t.Setenv("zero", "000")
	t.Setenv("string", "golang")
	t.Setenv("emptystring", "")
	t.Setenv("true", "true")
	t.Setenv("false", "false")

	cases := []testCase[string, bool]{
		{"num", "Numeric values should be truthy", func(t *testing.T, result bool, err error) bool {
			if result != true {
				t.Errorf("Expected %t, got %t", true, result)
				return false
			}
			return true
		}},
		{"zero", "Numeric value of \"0\" should be falsey", func(t *testing.T, result bool, err error) bool {
			if result != false {
				t.Errorf("Expected %t, got %t", false, result)
				return false
			}
			return true
		}},
		{"string", "Non empty strings should be truthy", func(t *testing.T, result bool, err error) bool {
			if result != true {
				t.Errorf("Expected %t, got %t", true, result)
				return false
			}
			return true
		}},
		{"emptystring", "Empty strings should be falsey", func(t *testing.T, result bool, err error) bool {
			if result != false {
				t.Errorf("Expected %t, got %t", false, result)
				return false
			}
			return true
		}},
		{"true", "Boolean values are parsed as booleans (true)", func(t *testing.T, result bool, err error) bool {
			if result != true {
				t.Errorf("Expected %t, got %t", true, result)
				return false
			}
			return true
		}},
		{"false", "Boolean values are parsed as booleans (false)", func(t *testing.T, result bool, err error) bool {
			if result != false {
				t.Errorf("Expected %t, got %t", false, result)
				return false
			}
			return true
		}},
	}

	runTestCases(t, "Bool()", ayika.Bool, cases)
}
