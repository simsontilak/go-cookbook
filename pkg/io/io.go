package io

import (
	"fmt"
	"strings"
)

func DisplayMainHeading(heading string) {
	displayHeading(heading, "=")
}

func DisplaySubHeading(heading string) {
	displayHeading(heading, "-")
}

func LineSeparator() {
	fmt.Printf("\n%s\n", strings.Repeat("#", 40))
}

func displayHeading(heading string, decoratorChar string) {
	sep := strings.Repeat(decoratorChar, len(heading))
	fmt.Printf("\n\n%s\n%s\n%s\n", sep, heading, sep)
}
