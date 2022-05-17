package ayika

import (
	"fmt"

	"github.com/fatih/color"
)

func reportError(err error, message string) {
	canShowErrors := Bool("AYIKA_VERBOSE")
	if err != nil && canShowErrors {
		fmt.Printf("%s: %s\n", color.HiRedString("Error"), message)
	}
}
