package datastructures

import "fmt"

func DemoSliceOp() {
	var x []int

	x = append(x, 1, 2, 3, 4)

	y := append(x, 3, 4, 5, 6)

	fmt.Println("Contents of x", x)

	fmt.Println("Contents of y", y)

	x = append(y, x...)

	fmt.Println("Contents of x after appending y", x)

}
