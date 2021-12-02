package main

import (
	"fmt"
	"github.com/eunicebjm/adventOfCode2021/rdr"
	"log"
)

func sonarSweeper(depths []int) int {
	count := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}

	return count
}

func main() {
	lines, err := rdr.ReadLines("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	intArr := rdr.ToIntArray(lines)
	count := sonarSweeper(intArr)
	fmt.Println(count)
}
