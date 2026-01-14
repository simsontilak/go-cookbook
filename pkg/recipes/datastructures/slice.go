package datastructures

import "fmt"

func DemoSliceOp() {
	var x = make([]int, 0, 10)

	x = append(x, 1, 2, 3, 4)

	y := append(x, 3, 4, 5, 6)

	fmt.Println("Contents of x", x, "Capacity of x", cap(x))

	fmt.Println("Contents of y", y, "Capacity of y", cap(y))

	x = append(y, x...)

	fmt.Println("Contents of x after appending y", x, "Capacity of x", cap(x))

}

func DemoSlicingSlice() {
	var x = make([]int, 0, 15)

	for i := 0; i < cap(x); i++ {
		x = append(x, i+1)
	}

	fmt.Println("Value of x = ", x)
	fmt.Println("Slice of first 4 elements", x[:4])
	fmt.Println("Slice of last 4 elements", x[len(x)-4:])
	fmt.Println("Slice of 6th element to 10th", x[5:10])

}
