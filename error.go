package ayika

import (
	"fmt"

	"github.com/fatih/color"
)

func reportError(err error, message string) {
	canShowErrors := Bool("AYIKA_SHOW_ERRORS")
	if err != nil && canShowErrors {
		fmt.Printf("%s: %s\n", color.HiRedString("Error"), message)
	}
}
