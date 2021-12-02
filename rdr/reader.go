package rdr

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ReadLines returns lines from a text file as a slice
func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return lines, nil
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ToIntArray(str []string) []int {
	var nums []int
	for idx := range str {
		i := StringToInt(str[idx])
		nums = append(nums, i)
	}

	return nums
}
