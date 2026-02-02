package random

import (
	"math"
	"math/rand"
)

func RandomRange(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func RandomRangeSlice(count int, min int, max int) []int {
	result := make([]int, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, RandomRange(min, max))
	}

	return result
}

func RandomRangeNormal(min int, max int) int {

	for {
		normalFloat := rand.NormFloat64()

		mean := float64(min+max) / 2.0
		std := 2.5

		scaled := normalFloat * std + mean

		result := int(math.Round(scaled))

		if result >= min && result <= max {
			return result
		}

	}
}

func RandomRangeSliceNormal(count int, min int, max int) []int {
	result := make([]int, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, RandomRangeNormal(min, max))
	}

	return result
}
