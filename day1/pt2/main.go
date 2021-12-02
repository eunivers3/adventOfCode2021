package main

import (
	"fmt"
	"github.com/eunicebjm/adventOfCode2021/rdr"
	"log"
)

func slidingWindow(depths []int) int {
	var (
		previous = depths[0] + depths[1] + depths[2]
		count    = 0
	)

	for i := 2; i < len(depths)-1; i++ {
		threeSum := depths[i] + depths[i-1] + depths[i+1]
		if threeSum > previous {
			count++
		}
		previous = threeSum

	}

	return count
}

func main() {
	lines, err := rdr.ReadLines("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	intArr := rdr.ToIntArray(lines)
	count := slidingWindow(intArr)
	fmt.Println(count)
}
