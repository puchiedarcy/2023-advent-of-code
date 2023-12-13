package main

import (
	"2023-advent-of-code/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func RunDay12() {
	file, err := os.Open("./inputs/day12input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalArrangements := 0
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		lineText := scanner.Text()
		halves := strings.Split(lineText, " ")
		springMap := strings.Repeat(halves[0]+"?", 5)
		springMap = springMap[:len(springMap)-1]
		damagedLengths := []int{}
		for i := 0; i < 5; i++ {
			for _, v := range helpers.ParseInts(halves[1]) {
				damagedLengths = append(damagedLengths, v)
			}
		}
		start := time.Now()
		matches := findMatches(springMap, damagedLengths, "", len(springMap))
		fmt.Println("Line", lineNumber, "done in:", time.Now().Sub(start))
		fmt.Println("Line", lineNumber, "matches:", matches)

		lineNumber++
		totalArrangements += matches

	}
	fmt.Println(totalArrangements)
}

func findMatches(springMap string, damagedLengths []int, mapInProgress string, mapLen int) int {
	matches := 0

	mandatoryChars := 0
	for _, v := range damagedLengths {
		mandatoryChars += v
	}

	availableSpots := strings.Count(springMap, "?")
	availableSpots += strings.Count(springMap, "#")

	if mandatoryChars > availableSpots {
		return 0
	}

	for i := 0; i < len(springMap); i++ {
		if springMap[i] == '.' {
			mapInProgress += "."
			continue
		}

		matches += applyDamageSet(springMap[i:], damagedLengths, mapInProgress, mapLen)
		if springMap[i] == '?' {
			mapInProgress += "."
		}
	}
	if len(damagedLengths) == 0 {
		if len(mapInProgress) == mapLen {
			return 1
		}
	}
	return matches
}

func applyDamageSet(springMap string, damagedLengths []int, mapInProgress string, mapLen int) int {
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

	return findMatches(nextSpringMap, nextDamagedLengths, mapInProgress, mapLen)
}
