package main

import (
	"log"
	"os"
	"testing"
)

func Test_Part1(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := readInput(file)
	got1 := Part1(input)
	want1 := 11605

	if got1 != want1 {
		t.Errorf("got %d, wanted %d", got1, want1)
	}
}

func Test_Part2(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := readInput(file)
	got2 := Part2(input)
	want2 := 12725

	if got2 != want2 {
		t.Errorf("got %d, wanted %d", got2, want2)
	}
}
