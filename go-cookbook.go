package main

import (
	"simsonlive.com/go-cookbook/pkg/io"
	"simsonlive.com/go-cookbook/pkg/recipes/datastructures"
	"simsonlive.com/go-cookbook/pkg/recipes/typemgmt"
)

func main() {
	type demoFunctions func()
	givenFunctions := []demoFunctions{
		typemgmt.DemoTypeConversion,
		typemgmt.DemoComplex,
		datastructures.ArrayPrintDemo,
		datastructures.DemoSliceOp,
		datastructures.DemoSlicingSlice,
		datastructures.DemoMapUniform,
		datastructures.DemoMapNormal,
	}

	io.DisplayMainHeading("Cookbook Demo")
	for _, operation := range givenFunctions {
		io.DisplayRunHeading(operation)
		operation()
		io.LineSeparator()
	}
}
