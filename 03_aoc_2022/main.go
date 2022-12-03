package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func calc_priority(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 65 + 27
	} else {
		return int(r) - 96
	}
}

func search_common_item(first string, second string) rune {
	for _, letter := range first {
		for _, letter2 := range second {
			if letter == letter2 {
				return letter
			}
		}
	}
	return 0
}

func search_common_item2(first string, second string, third string) rune {
	letters_map1 := make(map[rune]int)
	letters_map2 := make(map[rune]int)
	letters_map3 := make(map[rune]int)

	for _, letter := range first {
		letters_map1[letter]++
	}
	for _, letter := range second {
		letters_map2[letter]++
	}
	for _, letter := range third {
		letters_map3[letter]++
	}

	for key, _ := range letters_map1 {
		if _, ok2 := letters_map2[key]; ok2 {
			if _, ok3 := letters_map3[key]; ok3 {
				return rune(key)
			}
		}
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		first_compartment := line[0 : len(line)/2]
		second_compartment := line[len(line)/2:]
		common_item := search_common_item(first_compartment, second_compartment)
		priority := calc_priority(common_item)
		sum += priority
	}

	fmt.Printf("Day three aoc 2022\n")
	fmt.Printf("Part one: %d\n", sum)

	sum2 := 0
	for i := 0; i < len(lines); i += 3 {
		first_elves := lines[i]
		second_elves := lines[i+1]
		third_elves := lines[i+2]
		common_item2 := search_common_item2(first_elves, second_elves, third_elves)
		priority := calc_priority(common_item2)
		sum2 += priority
	}
	fmt.Printf("Part two: %d\n", sum2)
}
