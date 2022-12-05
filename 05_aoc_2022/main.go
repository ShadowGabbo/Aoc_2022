package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return s[:l-1], s[l-1]
}

func move(n_stack string, from_string string, to_string string, stacks []stack) []stack {
	limit, _ := strconv.Atoi(n_stack)
	from, _ := strconv.Atoi(from_string)
	to, _ := strconv.Atoi(to_string)
	for i := 0; i < limit; i++ {
		pop := ""
		stacks[from-1], pop = stacks[from-1].Pop()
		stacks[to-1] = stacks[to-1].Push(pop)
	}
	return stacks
}

func move2(n_stack string, from_string string, to_string string, stacks []stack) []stack {
	limit, _ := strconv.Atoi(n_stack)
	from, _ := strconv.Atoi(from_string)
	to, _ := strconv.Atoi(to_string)

	slice_pop := []string{}
	for i := 0; i < limit; i++ {
		pop := ""
		stacks[from-1], pop = stacks[from-1].Pop()
		slice_pop = append([]string{pop}, slice_pop...)
		if i == limit-1 {
			for _, pop := range slice_pop {
				stacks[to-1] = stacks[to-1].Push(pop)
			}
		}
	}
	return stacks
}

func inizialize() []stack {
	// stacks
	s1 := make(stack, 0)
	s2 := make(stack, 0)
	s3 := make(stack, 0)
	s4 := make(stack, 0)
	s5 := make(stack, 0)
	s6 := make(stack, 0)
	s7 := make(stack, 0)
	s8 := make(stack, 0)
	s9 := make(stack, 0)
	s1 = s1.Push("D")
	s1 = s1.Push("T")
	s1 = s1.Push("R")
	s1 = s1.Push("B")
	s1 = s1.Push("J")
	s1 = s1.Push("L")
	s1 = s1.Push("W")
	s1 = s1.Push("G")
	s2 = s2.Push("S")
	s2 = s2.Push("W")
	s2 = s2.Push("C")
	s3 = s3.Push("R")
	s3 = s3.Push("Z")
	s3 = s3.Push("T")
	s3 = s3.Push("M")
	s4 = s4.Push("D")
	s4 = s4.Push("T")
	s4 = s4.Push("C")
	s4 = s4.Push("H")
	s4 = s4.Push("S")
	s4 = s4.Push("P")
	s4 = s4.Push("V")
	s5 = s5.Push("G")
	s5 = s5.Push("P")
	s5 = s5.Push("T")
	s5 = s5.Push("L")
	s5 = s5.Push("D")
	s5 = s5.Push("Z")
	s6 = s6.Push("F")
	s6 = s6.Push("B")
	s6 = s6.Push("R")
	s6 = s6.Push("Z")
	s6 = s6.Push("J")
	s6 = s6.Push("Q")
	s6 = s6.Push("C")
	s6 = s6.Push("D")
	s7 = s7.Push("S")
	s7 = s7.Push("B")
	s7 = s7.Push("D")
	s7 = s7.Push("J")
	s7 = s7.Push("M")
	s7 = s7.Push("F")
	s7 = s7.Push("T")
	s7 = s7.Push("R")
	s8 = s8.Push("L")
	s8 = s8.Push("H")
	s8 = s8.Push("R")
	s8 = s8.Push("B")
	s8 = s8.Push("T")
	s8 = s8.Push("V")
	s8 = s8.Push("M")
	s9 = s9.Push("Q")
	s9 = s9.Push("P")
	s9 = s9.Push("D")
	s9 = s9.Push("S")
	s9 = s9.Push("V")

	// inizialize all stacks
	stacks := make([]stack, 9)
	stacks[0] = s1
	stacks[1] = s2
	stacks[2] = s3
	stacks[3] = s4
	stacks[4] = s5
	stacks[5] = s6
	stacks[6] = s7
	stacks[7] = s8
	stacks[8] = s9
	return stacks
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stacks := inizialize()

	// part1
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		line_splitted := strings.Split(line, " ")
		n_stack := line_splitted[1]
		from := line_splitted[3]
		to := line_splitted[5]
		stacks = move(n_stack, from, to, stacks)
	}

	end_tops := ""
	for i := 0; i < 9; i++ {
		len := len(stacks[i])
		end := stacks[i][len-1]
		end_tops += end
	}

	fmt.Printf("Day five aoc 2022\n")
	fmt.Printf("Part one: %s\n", end_tops)

	// part two
	stacks2 := inizialize()

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		line_splitted := strings.Split(line, " ")
		n_stack := line_splitted[1]
		from := line_splitted[3]
		to := line_splitted[5]

		stacks2 = move2(n_stack, from, to, stacks2)
	}

	end_tops2 := ""
	for i := 0; i < 9; i++ {
		len := len(stacks2[i])
		end := stacks2[i][len-1]
		end_tops2 += end
	}
	fmt.Printf("Part two: %s\n", end_tops2)
}
