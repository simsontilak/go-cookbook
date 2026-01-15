package io

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

const LineWidth = 150

func DisplayMainHeading(heading string) {
	displayHeading(heading, "=")
}

func DisplaySubHeading(heading string) {
	displayHeading(heading, "-")
}

func DisplayRunHeading(i interface{}) {
	functionName := GetFunctionName(i)
	DisplaySubHeading("Running " + functionName + " ...")
}

func LineSeparator() {
	fmt.Printf("\n%s\n", strings.Repeat("-", LineWidth))
}

func displayHeading(heading string, decoratorChar string) {
	sep := strings.Repeat(decoratorChar, LineWidth)
	fmt.Printf("\n\n%s\n%s\n%s\n", sep, heading, sep)
}

func GetFunctionName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	// Split the name by the dot delimiter to get the last part (the function name)
	splits := strings.Split(fullName, ".")
	return splits[len(splits)-1]
}
