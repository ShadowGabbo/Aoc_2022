package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Play rune

type turn struct {
	my_choice       Play
	opponent_choice Play
}

const (
	Rock    Play = 'A'
	Paper        = 'B'
	Scissor      = 'C'
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// part one
	scanner := bufio.NewScanner(file)
	total_points := 0
	rock_paper_scissor := map[turn]int{
		{my_choice: 'A', opponent_choice: 'A'}: 4,
		{my_choice: 'B', opponent_choice: 'B'}: 5,
		{my_choice: 'C', opponent_choice: 'C'}: 6,
		{my_choice: 'A', opponent_choice: 'B'}: 1,
		{my_choice: 'A', opponent_choice: 'C'}: 7,
		{my_choice: 'A', opponent_choice: 'A'}: 4,
		{my_choice: 'B', opponent_choice: 'A'}: 8,
		{my_choice: 'B', opponent_choice: 'C'}: 2,
		{my_choice: 'C', opponent_choice: 'A'}: 3,
		{my_choice: 'C', opponent_choice: 'B'}: 9}
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		opponent_choice := line[0]
		my_choice := line[2] - 23 // so i can use rune as numbers
		turn := turn{my_choice: Play(my_choice), opponent_choice: Play(opponent_choice)}
		total_points += rock_paper_scissor[turn]
	}
	fmt.Printf("Day two aoc 2022\n")
	fmt.Printf("Part one: %d\n", total_points)

	//part two
	total_points2 := 0
	rock_paper_scissor2 := map[turn]rune{
		{my_choice: 'A', opponent_choice: 'A'}: 'C',
		{my_choice: 'B', opponent_choice: 'B'}: 'B',
		{my_choice: 'C', opponent_choice: 'C'}: 'A',
		{my_choice: 'A', opponent_choice: 'B'}: 'A',
		{my_choice: 'A', opponent_choice: 'C'}: 'B',
		{my_choice: 'B', opponent_choice: 'A'}: 'A',
		{my_choice: 'B', opponent_choice: 'C'}: 'C',
		{my_choice: 'C', opponent_choice: 'A'}: 'B',
		{my_choice: 'C', opponent_choice: 'B'}: 'C'}

	for _, line := range lines {
		opponent_choice2 := line[0]
		condition := line[2] - 23 // so i can use rune as numbers
		get_choice := turn{my_choice: Play(opponent_choice2), opponent_choice: Play(condition)}
		my_choice2 := rock_paper_scissor2[get_choice]
		turn := turn{my_choice: Play(my_choice2), opponent_choice: Play(opponent_choice2)}
		total_points2 += rock_paper_scissor[turn]
	}
	fmt.Printf("Part two: %d\n", total_points2)

}
