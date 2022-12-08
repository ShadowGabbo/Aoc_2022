package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func count_lines() ([]string, int) {
	var lines []string
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// part1
	scanner := bufio.NewScanner(file)
	n_lines := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		n_lines++
	}
	return lines, n_lines
}

func part1(trees [][]int, n_lines int) int {
	counter := 0

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			if i == 0 || j == 0 || i == len(trees)-1 || j == len(trees[i])-1 {
				counter++
			} else {
				if is_visible(trees, i, j, n_lines) {
					counter++
				}
			}
		}
	}

	return counter
}

func part2(trees [][]int, n_lines int) int {
	max_scenic_score := 0

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			if i == 0 || j == 0 || i == len(trees)-1 || j == len(trees[i])-1 {
				continue
			} else {
				if scenic_score := get_scenic_score(trees, i, j, n_lines); scenic_score > max_scenic_score {
					max_scenic_score = scenic_score
				}
			}
		}
	}

	return max_scenic_score
}

func get_scenic_score(trees [][]int, i int, j int, n_lines int) int {
	num := trees[i][j]
	top, down, left, right := get_seq(trees, i, j, n_lines)
	viewing_distance1 := calc_distance(num, top)
	viewing_distance2 := calc_distance(num, down)
	viewing_distance3 := calc_distance(num, left)
	viewing_distance4 := calc_distance(num, right)
	return viewing_distance1 * viewing_distance2 * viewing_distance3 * viewing_distance4
}

func calc_distance(num int, seq []int) int {
	distance := 0
	for _, num2 := range seq {
		if num2 >= num {
			distance++
			break
		}
		distance++
	}
	return distance
}

func is_visible(trees [][]int, i int, j int, n_lines int) bool {
	num := trees[i][j]
	top, down, left, right := get_seq(trees, i, j, n_lines)

	flag1 := check_sequence(num, top)
	flag2 := check_sequence(num, down)
	flag3 := check_sequence(num, right)
	flag4 := check_sequence(num, left)

	return flag1 || flag2 || flag3 || flag4
}

func check_sequence(num int, seq []int) bool {
	for _, num2 := range seq {
		if num2 >= num {
			return false
		}
	}
	return true
}

func get_seq(trees [][]int, i int, j int, n_lines int) ([]int, []int, []int, []int) {
	top := []int{}
	down := []int{}
	left := []int{}
	right := []int{}

	for k := i - 1; k >= 0; k-- {
		top = append(top, trees[k][j])
	}
	for k := i + 1; k < n_lines; k++ {
		down = append(down, trees[k][j])
	}
	for k := j - 1; k >= 0; k-- {
		left = append(left, trees[i][k])
	}
	for k := j + 1; k < len(trees[i]); k++ {
		right = append(right, trees[i][k])
	}
	return top, down, left, right
}

func main() {
	lines, n_lines := count_lines()
	var trees [][]int = make([][]int, n_lines)
	for i := 0; i < n_lines; i++ {
		line := lines[i]
		trees[i] = make([]int, len(line))
		for j, v := range line {
			num := int(v - '0')
			trees[i][j] = num
		}
	}

	fmt.Printf("Day eight aoc 2022\n")
	fmt.Printf("Part one: %d\n", part1(trees, n_lines))
	fmt.Printf("Part two: %d\n", part2(trees, n_lines))
}
