package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

	scanner := bufio.NewScanner(file)
	somma := 0
	map_elves := make(map[int]int)
	counter := 0
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if num != 0 {
			somma += num
		} else {
			map_elves[counter] = somma
			somma = 0
			counter++
		}
	}

	// sort map
	var ss []kv
	for k, v := range map_elves {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	fmt.Printf("Day one aoc 2022\n")
	fmt.Printf("Part one: %d\n", ss[0].Value)
	fmt.Printf("Part two: %d\n", (ss[0].Value + ss[1].Value + ss[2].Value))
}
