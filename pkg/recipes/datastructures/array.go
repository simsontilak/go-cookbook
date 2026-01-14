package datastructures

import "fmt"

func ArrayPrintDemo() {
	var data [15]int

	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	for j := 0; j < len(data); j++ {
		fmt.Printf("%d %d %d\n", j, data[j], data[j]*10)
	}
}
