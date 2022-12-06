package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func is_marker(occurrency map[string]int) bool {
	for _, v := range occurrency {
		if v > 1 {
			return false
		}
	}
	return true
}

func make_map(line string, i int, occ map[string]int) map[string]int {
	for j := 0; j < 14; j++ {
		occ[string(line[i+j])]++
	}
	return occ
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// part1
	scanner := bufio.NewScanner(file)
	line := ""
	if scanner.Scan() {
		line = scanner.Text()
	}

	characters := 0
	for i := 0; i < len(line); i++ {
		occurrency := make(map[string]int)
		letter1 := line[i]
		letter2 := line[i+1]
		letter3 := line[i+2]
		letter4 := line[i+3]
		occurrency[string(letter1)]++
		occurrency[string(letter2)]++
		occurrency[string(letter3)]++
		occurrency[string(letter4)]++
		marker := is_marker(occurrency)
		if marker {
			characters = i + 3 + 1 // cause start from 0
			break
		}
	}

	fmt.Printf("Day six aoc 2022\n")
	fmt.Printf("Part one: %d\n", characters)

	// part two
	characters2 := 0
	for i := 0; i < len(line); i++ {
		occurrency := make(map[string]int)
		occurrency = make_map(line, i, occurrency)
		marker := is_marker(occurrency)
		if marker {
			characters2 = i + 14
			break
		}
	}
	fmt.Printf("Part two: %d\n", characters2)
}
