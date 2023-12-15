package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func RunDay15() {
	file, err := os.Open("./inputs/day15input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalHASH := 0
	currentValue := 0
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lineText := scanner.Text()
	steps := strings.Split(lineText, ",")

	for _, step := range steps {
		for _, c := range step {
			currentValue += int(c)
			currentValue *= 17
			currentValue %= 256
		}
		totalHASH += currentValue
		currentValue = 0
	}
	fmt.Println(totalHASH)
}
