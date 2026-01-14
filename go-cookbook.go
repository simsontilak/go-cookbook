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
}
