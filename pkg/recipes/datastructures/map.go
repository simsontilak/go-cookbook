package datastructures

import (
	"fmt"
	"strings"

	"simsonlive.com/go-cookbook/pkg/random"
)

func DemoMapUniform() {
	//generate 1000 random numbers between 1 to 10
	aSlice := random.RandomRangeSlice(500, 1, 10)
	demoMapFrequencyDistribution(aSlice)
}

func DemoMapNormal() {
	//generate 1000 random numbers between 1 to 10
	aSlice := random.RandomRangeSliceNormal(500, 1, 10)
	demoMapFrequencyDistribution(aSlice)
}

func demoMapFrequencyDistribution(aSlice []int) {
	frequencyMap := map[int]int{}

	for i := 0; i < len(aSlice); i++ {
		val, ok := frequencyMap[aSlice[i]]
		if ok {
			frequencyMap[aSlice[i]] = val + 1
		} else {
			frequencyMap[aSlice[i]] = 1
		}
	}

	fmt.Println("Frequency Distribution: ")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%4d   %-90s  %4d\n", i, strings.Repeat("|", frequencyMap[i]), frequencyMap[i])
	}
}
