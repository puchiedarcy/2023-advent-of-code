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
	file, err := os.Open("./inputs/day12inputsmall.txt")
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

		fmt.Println(springMap, damagedLengths)

		findMatches(springMap, damagedLengths, "")

	}
	fmt.Println(totalArrangements)
}

func findMatches(springMap string, damagedLengths []int, mapInProgress string) {
	for i := 0; i < len(springMap); i++ {
		if springMap[i] == '.' {
			continue
		} else {
			mapInProgress += "."
		}

		applyDamageSet(springMap[i:], damagedLengths, mapInProgress)
	}
	if len(damagedLengths) == 0 && len(springMap) == 0 {
		fmt.Println("MATCH:", mapInProgress)
	}
}

func applyDamageSet(springMap string, damagedLengths []int, mapInProgress string) {
	for i := 0; i < damagedLengths[0]; i++ {
		if i >= len(springMap) {
			return
		}
		currentChar := springMap[i]
		if currentChar == '.' {
			return
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
			return
		}
		nextSpringMap = nextSpringMap[1:]
	}

	for len(nextSpringMap) > 0 && nextSpringMap[0] == '.' {
		nextSpringMap = nextSpringMap[1:]
		mapInProgress += "."
	}

	findMatches(nextSpringMap, nextDamagedLengths, mapInProgress)
}
