package main

import (
	"fmt"
	"github.com/eunicebjm/adventOfCode2021/rdr"
	"log"
	"strings"
)

func submarine(cmds []string) int {
	var (
		aim     int
		forward int
		depth   int
	)

	for i := range cmds {
		arr := strings.Split(cmds[i], " ")
		magnitude := rdr.StringToInt(arr[1])
		switch arr[0] {
		case "forward":
			forward += magnitude
			depth += aim * magnitude
		case "up":
			aim -= magnitude
		case "down":
			aim += magnitude
		default:
			break
		}
	}

	return forward * depth
}

func main() {
	lines, err := rdr.ReadLines("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	count := submarine(lines)
	fmt.Println(count)
}
