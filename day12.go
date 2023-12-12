package main

import (
	"2023-advent-of-code/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func RunDay12() {
	file, err := os.Open("./inputs/day12input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalArrangements := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		halves := strings.Split(lineText, " ")
		springMap := halves[0]
		damagedLengths := helpers.ParseInts(halves[1])

		matches := findMatches(springMap, damagedLengths, "")
		totalArrangements += matches
	}
	fmt.Println(totalArrangements)
}

func findMatches(springMap string, damagedLengths []int, mapInProgress string) int {
	matches := 0
	for i := 0; i < len(springMap); i++ {
		if springMap[i] == '.' {
			mapInProgress += "."
			continue
		}

		matches += applyDamageSet(springMap[i:], damagedLengths, mapInProgress)
		if springMap[i] == '?' {
			mapInProgress += "."
		}
	}
	if len(damagedLengths) == 0 && len(springMap) == 0 {
		return 1
	}
	return matches
}

func applyDamageSet(springMap string, damagedLengths []int, mapInProgress string) int {
	if len(damagedLengths) == 0 {
		return 0
	}

	for i := 0; i < damagedLengths[0]; i++ {
		if i >= len(springMap) {
			return 0
		}
		currentChar := springMap[i]
		if currentChar == '.' {
			return 0
		} else {
			mapInProgress += "#"
		}
	}

	prevDamageLength := damagedLengths[0]
	nextSpringMap := springMap[prevDamageLength:]
	nextDamagedLengths := damagedLengths[1:]
	if len(nextDamagedLengths) != 0 {
		mapInProgress += "."
		if len(nextSpringMap) == 0 {
			return 0
		}
		if nextSpringMap[0] == '#' {
			return 0
		}
		nextSpringMap = nextSpringMap[1:]
	}

	for len(nextSpringMap) > 0 && nextSpringMap[0] == '.' {
		nextSpringMap = nextSpringMap[1:]
		mapInProgress += "."
	}

	return findMatches(nextSpringMap, nextDamagedLengths, mapInProgress)
}
