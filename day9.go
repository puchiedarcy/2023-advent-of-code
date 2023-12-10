package main

import (
	"2023-advent-of-code/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunDay9Part1() {
	file, err := os.Open("./inputs/day9input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sumValues := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		numbers := helpers.ParseInts(lineText)
		nextNumber := CalculateNextNumber(numbers)
		sumValues += nextNumber
	}

	fmt.Println(sumValues)
}

func CalculateNextNumber(currentLine []int) int {
	nextLine := []int{}
	allZeros := true
	for i := 1; i < len(currentLine); i++ {
		nextNumber := currentLine[i] - currentLine[i-1]
		nextLine = append(nextLine, nextNumber)
		if nextNumber != 0 {
			allZeros = false
		}
	}

	if allZeros {
		return currentLine[len(currentLine)-1]
	}

	return currentLine[len(currentLine)-1] + CalculateNextNumber(nextLine)
}
