package main

import (
	"fmt"
	"github.com/eunicebjm/adventOfCode2021/rdr"
	"log"
	"strings"
)

func submarine(cmds []string) int {
	var (
		forward int
		depth   int
	)

	for i := range cmds {
		arr := strings.Split(cmds[i], " ")
		switch arr[0] {
		case "forward":
			forward += rdr.StringToInt(arr[1])
		case "up":
			depth -= rdr.StringToInt(arr[1])
		case "down":
			depth += rdr.StringToInt(arr[1])
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
