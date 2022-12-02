package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type kv struct {
	Key   int
	Value int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// part one
	scanner := bufio.NewScanner(file)
	total_points := 0
	rock_paper_scissor := map[string]int{"AA": 4, "BB": 5, "CC": 6, "AB": 1, "AC": 7, "BA": 8, "BC": 2, "CA": 3, "CB": 9}
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		opponent_choice := line[0]
		my_choice := line[2] - 23 // so i can use rune as numbers
		total_points += rock_paper_scissor[string(my_choice)+string(opponent_choice)]
	}
	fmt.Printf("Day two aoc 2022\n")
	fmt.Printf("Part one: %d\n", total_points)

	// part two
	total_points2 := 0
	rock_paper_scissor2 := map[string]string{"AA": "C", "BB": "B", "CC": "A", "AB": "A", "AC": "B", "BA": "A", "BC": "C", "CA": "B", "CB": "C"}

	for _, line := range lines {
		opponent_choice2 := line[0]
		condition := line[2] - 23 // so i can use rune as numbers
		my_choice2 := rock_paper_scissor2[string(opponent_choice2)+string(condition)]
		total_points2 += rock_paper_scissor[string(my_choice2)+string(opponent_choice2)]
	}
	fmt.Printf("Part two: %d\n", total_points2)
}
