package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Printf("\n========================\n")
	fmt.Printf("\nHere is a list of Demos:\n")
	fmt.Printf("\n========================\n")
	demoComplex()
	demoTypeConversion()
}

func demoTypeConversion() {
	fmt.Printf("\nDemo of type conversion: \n")
	var x int = 10
	var y float64 = 44.5324
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println("Sum1 = ", sum1, "Sum2 = ", sum2)
	var a int = 10
	var b byte = 20
	var sum3 int = a + int(b)
	var sum4 byte = byte(a) + b
	fmt.Println("Sum3 = ", sum3, "Sum4 = ", sum4)
}

func demoComplex() {
	fmt.Printf("\nDemo of complex numbers: \n")
	first := complex(4.2, 3.3)
	second := complex(5.3, 2.2)

	sum := first + second
	prod := first * second
	div := first / second
	diff := first - second
	real := real(first)
	img := imag(second)
	absVal := cmplx.Abs(first)

	fmt.Printf("\nSum: %4.2f\nProduct: %4.2f\nDivision: %4.2f"+
		"\nDifference: %4.2f\nReal: %4.2f\nImaginary: %4.2f"+
		"\nAbsolute Value: %4.2f\n", sum, prod, div, diff, real, img, absVal)
}
