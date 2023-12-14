package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunDay14() {
	file, err := os.Open("./inputs/day14input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	platform := make(map[int]string)
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		platform[lineNumber] = lineText
		lineNumber++
	}

	total := 0
	for col := 0; col < len(platform); col++ {
		colTotal := 0
		value := len(platform)
		for row := 0; row < len(platform); row++ {
			c := platform[row][col]
			if c == 'O' {
				colTotal += value
				value--
			} else if c == '.' {
				continue
			} else if c == '#' {
				value = len(platform) - row - 1
			}
		}
		total += colTotal
	}

	fmt.Println(total)
}
