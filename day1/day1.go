package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunDay1Part1() {
	file, err := os.Open("./day1/day1input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	sum := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		leftIndex := 0
		leftDigit := 0

		for leftIndex < len(lineText) {
			asciiDigit := lineText[leftIndex]
			if asciiDigit >= 48 && asciiDigit <= 57 {
				leftDigit = int(asciiDigit) - 48
				break
			}
			leftIndex++
		}

		rightIndex := len(lineText) - 1
		rightDigit := 0
		for rightIndex >= 0 {
			asciiDigit := lineText[rightIndex]
			if asciiDigit >= 48 && asciiDigit <= 57 {
				rightDigit = int(asciiDigit) - 48
				break
			}
			rightIndex--
		}

		sum += leftDigit*10 + rightDigit
	}
	fmt.Println(sum)
}
