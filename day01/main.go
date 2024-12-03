package main

import (
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	helper "advent-of-code-2024/helper"
)

//go:embed inputFiles/*
var embeddedFiles embed.FS

func part1(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	if len(left) != len(right) {
		panic("left and right should have the same length")
	}
	distance := 0
	for i := 0; i < len(left); i++ {
		distance += helper.Abs(right[i] - left[i])
	}
	return distance
}

func part2(left, right []int) int {
	rightMap := make(map[int]int)
	for _, r := range right {
		rightMap[r] += 1
	}
	similarity := 0
	for _, l := range left {
		similarity += l * rightMap[l]
	}
	return similarity
}

func main() {
	left := []int{}
	right := []int{}
	for _, line := range helper.ReadFile(embeddedFiles, "inputFiles/part1.txt") {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		left = append(left, helper.Must(strconv.Atoi(parts[0])))
		right = append(right, helper.Must(strconv.Atoi(parts[1])))
	}
	fmt.Printf("Day 01, Part 01: %d\n", part1(left, right))
	fmt.Printf("Day 01, Part 02: %d\n", part2(left, right))
}
