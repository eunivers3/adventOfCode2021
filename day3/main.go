package main

import (
	"fmt"
	"github.com/eunicebjm/adventOfCode2021/day3/input"
	"strconv"
)

func binaryDiagnosticPt1(nums []string) int64 {
	countOnes := make([]int, len(nums[0]))
	for _, bits := range nums {
		for i := range bits {
			if string(bits[i]) == "1" {
				countOnes[i] += 1
			}
		}
	}

	return getPower(countOnes, len(nums))
}

func isOne(onesCount, total int) bool {
	return float32(onesCount)/float32(total) > 0.5
}

func binaryToInt(b string) int64 {
	if i, err := strconv.ParseInt(b, 2, 64); err != nil {
		fmt.Println("error converting binaryToInt:", b)
		return 0
	} else {
		return i
	}
}

func getPower(arr []int, total int) int64 {
	var (
		gamma, epsilon string
	)

	for i := range arr {
		if isOne(arr[i], total) {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	return binaryToInt(gamma) * binaryToInt(epsilon)
}

func main() {
	nums := input.Get(input.Input)
	result := binaryDiagnosticPt1(nums)
	fmt.Println(result)
}
