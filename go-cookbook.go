package main

import (
	"simsonlive.com/go-cookbook/pkg/io"
	"simsonlive.com/go-cookbook/pkg/recipes/datastructures"
	"simsonlive.com/go-cookbook/pkg/recipes/typemgmt"
)

func main() {
	io.DisplayMainHeading("Cookbook Demo")

	io.DisplaySubHeading("Type Conversion")
	typemgmt.DemoTypeConversion()
	io.LineSeparator()

	io.DisplaySubHeading("Complex Type")
	typemgmt.DemoComplex()
	io.LineSeparator()

	io.DisplayMainHeading("Arrays")
	datastructures.ArrayPrintDemo()
	io.LineSeparator()

	io.DisplayMainHeading("Slices")
	datastructures.DemoSliceOp()
	io.LineSeparator()

	io.DisplayMainHeading("Slices Operation")
	datastructures.DemoSlicingSlice()
	io.LineSeparator()

	io.DisplayMainHeading("Map Uniform Algorithm")
	datastructures.DemoMapUniform()
	io.LineSeparator()

	io.DisplayMainHeading("Map Normal Algorithm")
	datastructures.DemoMapNormal()
	io.LineSeparator()
}
