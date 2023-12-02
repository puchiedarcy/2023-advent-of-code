package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunDay1Part2() {
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
			} else {
				if len(lineText)-leftIndex >= 3 {
					threeLetterNumber := lineText[leftIndex : leftIndex+3]
					if threeLetterNumber == "one" {
						leftDigit = 1
						break
					} else if threeLetterNumber == "two" {
						leftDigit = 2
						break
					} else if threeLetterNumber == "six" {
						leftDigit = 6
						break
					}
				}

				if len(lineText)-leftIndex >= 4 {
					threeLetterNumber := lineText[leftIndex : leftIndex+4]
					if threeLetterNumber == "four" {
						leftDigit = 4
						break
					} else if threeLetterNumber == "five" {
						leftDigit = 5
						break
					} else if threeLetterNumber == "nine" {
						leftDigit = 9
						break
					}
				}

				if len(lineText)-leftIndex >= 5 {
					threeLetterNumber := lineText[leftIndex : leftIndex+5]
					if threeLetterNumber == "three" {
						leftDigit = 3
						break
					} else if threeLetterNumber == "seven" {
						leftDigit = 7
						break
					} else if threeLetterNumber == "eight" {
						leftDigit = 8
						break
					}
				}

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
			} else {
				if len(lineText)-rightIndex >= 3 {
					threeLetterNumber := lineText[rightIndex : rightIndex+3]
					if threeLetterNumber == "one" {
						rightDigit = 1
						break
					} else if threeLetterNumber == "two" {
						rightDigit = 2
						break
					} else if threeLetterNumber == "six" {
						rightDigit = 6
						break
					}
				}

				if len(lineText)-rightIndex >= 4 {
					threeLetterNumber := lineText[rightIndex : rightIndex+4]
					if threeLetterNumber == "four" {
						rightDigit = 4
						break
					} else if threeLetterNumber == "five" {
						rightDigit = 5
						break
					} else if threeLetterNumber == "nine" {
						rightDigit = 9
						break
					}
				}

				if len(lineText)-rightIndex >= 5 {
					threeLetterNumber := lineText[rightIndex : rightIndex+5]
					if threeLetterNumber == "three" {
						rightDigit = 3
						break
					} else if threeLetterNumber == "seven" {
						rightDigit = 7
						break
					} else if threeLetterNumber == "eight" {
						rightDigit = 8
						break
					}
				}

			}
			rightIndex--
		}

		sum += leftDigit*10 + rightDigit
	}
	fmt.Println(sum)
}

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
		fmt.Println(lineText, leftDigit, rightDigit, sum)
	}
	fmt.Println(sum)
}
