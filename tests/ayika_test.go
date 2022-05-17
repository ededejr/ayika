package ayika_test

import (
	"testing"

	"github.com/fatih/color"
)

var testWriter = color.New(color.Faint)

type testCase[Parameters any, Result any] struct {
	in          Parameters
	description string
	expected    func(t *testing.T, result Result, err error) bool
}

func runTestCasesWithErrors[Parameters any, Result any](t *testing.T, name string, f func(Parameters) (Result, error), cases []testCase[Parameters, Result]) {
	testWriter.Println(name)
	for _, test := range cases {
		printTestDescription(test.description)
		observed, err := f(test.in)
		success := test.expected(t, observed, err)
		printTestIcon(success)
	}
}

func runTestCases[Parameters any, Result any](t *testing.T, name string, f func(Parameters) Result, cases []testCase[Parameters, Result]) {
	testWriter.Println(name)
	for _, test := range cases {
		printTestDescription(test.description)
		observed := f(test.in)
		success := test.expected(t, observed, nil)
		printTestIcon(success)
	}
}

func printTestDescription(description string) {
	testWriter.Printf("  %s", description)
}

func printTestIcon(success bool) {
	if success {
		testWriter.Printf(color.HiGreenString(" ✔\n"))
	} else {
		testWriter.Printf(color.HiRedString(" 🅧\n"))
	}
}
