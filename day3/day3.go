package day3

import (
	"2023-advent-of-code/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func RunDay3Part2() {
	file, err := os.Open("./day3/day3input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	previousLine := scanner.Text()
	scanner.Scan()
	currentLine := scanner.Text()
	scanner.Scan()
	nextLine := scanner.Text()

	total := 0
	for nextLine != "" {
		for symbolIndex, c := range currentLine {
			char := string(c)
			if char == "." {
				continue
			}

			adjacentNumbers := []int{}
			if char == "*" {
				checkTopCenter := true
				checkTopRight := true
				topLeftChar := string(previousLine[symbolIndex-1])
				if helpers.IsDigit(topLeftChar) {
					if helpers.IsDigit(string(previousLine[symbolIndex])) {
						checkTopCenter = false
						if helpers.IsDigit(string(previousLine[symbolIndex+1])) {
							checkTopRight = false
						}
					}

					numberStack := helpers.Stack{}
					StackLeft(previousLine, symbolIndex-2, &numberStack)
					i, _ := strconv.Atoi(string(previousLine[symbolIndex-1]))
					numberStack.Push(i)
					StackRight(previousLine, symbolIndex, &numberStack)

					adjacentNumbers = append(adjacentNumbers, numberStack.ConvertToNumber())
				}

				if checkTopCenter {
					topCenterChar := string(previousLine[symbolIndex])
					if helpers.IsDigit(topCenterChar) {
						checkTopRight = false

						numberStack := helpers.Stack{}
						StackRight(previousLine, symbolIndex, &numberStack)

						adjacentNumbers = append(adjacentNumbers, numberStack.ConvertToNumber())
					}
				}

				if checkTopRight {
					topRightChar := string(previousLine[symbolIndex+1])
					if helpers.IsDigit(topRightChar) {
						numberStack := helpers.Stack{}
						StackRight(previousLine, symbolIndex+1, &numberStack)

						adjacentNumbers = append(adjacentNumbers, numberStack.ConvertToNumber())
					}
				}

				leftChar := string(currentLine[symbolIndex-1])
				if helpers.IsDigit(leftChar) {
					numberStack := helpers.Stack{}
					StackLeft(currentLine, symbolIndex-1, &numberStack)

					adjacentNumbers = append(adjacentNumbers, numberStack.ConvertToNumber())
				}

				rightChar := string(currentLine[symbolIndex+1])
				if helpers.IsDigit(rightChar) {
					numberStack := helpers.Stack{}
					StackRight(currentLine, symbolIndex+1, &numberStack)

					adjacentNumbers = append(adjacentNumbers, numberStack.ConvertToNumber())
				}

				checkBottomCenter := true
				checkBottomRight := true
				bottomLeftChar := string(nextLine[symbolIndex-1])
				if helpers.IsDigit(bottomLeftChar) {
					if helpers.IsDigit(string(nextLine[symbolIndex])) {
						checkBottomCenter = false
						if helpers.IsDigit(string(nextLine[symbolIndex+1])) {
							checkBottomRight = false
						}
					}

					numberStack := helpers.Stack{}
					StackLeft(nextLine, symbolIndex-2, &numberStack)
					i, _ := strconv.Atoi(string(nextLine[symbolIndex-1]))
					numberStack.Push(i)
					StackRight(nextLine, symbolIndex, &numberStack)

					adjacentNumbers = append(adjacentNumbers, numberStack.ConvertToNumber())
				}

				if checkBottomCenter {
					bottomCenterChar := string(nextLine[symbolIndex])
					if helpers.IsDigit(bottomCenterChar) {
						checkBottomRight = false

						numberStack := helpers.Stack{}
						StackRight(nextLine, symbolIndex, &numberStack)

						adjacentNumbers = append(adjacentNumbers, numberStack.ConvertToNumber())
					}
				}

				if checkBottomRight {
					bottomRightChar := string(nextLine[symbolIndex+1])
					if helpers.IsDigit(bottomRightChar) {
						numberStack := helpers.Stack{}
						StackRight(nextLine, symbolIndex+1, &numberStack)

						adjacentNumbers = append(adjacentNumbers, numberStack.ConvertToNumber())
					}
				}

				if len(adjacentNumbers) == 2 {
					ratio := adjacentNumbers[0] * adjacentNumbers[1]
					total += ratio
				}
			}
		}

		previousLine = currentLine
		currentLine = nextLine
		scanner.Scan()
		nextLine = scanner.Text()
		fmt.Println(total)
	}
}

func RunDay3Part1() {
	file, err := os.Open("./day3/day3input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	previousLine := scanner.Text()
	scanner.Scan()
	currentLine := scanner.Text()
	scanner.Scan()
	nextLine := scanner.Text()

	total := 0
	lineNumber := 1
	for nextLine != "" {
		fmt.Println(lineNumber, currentLine)
		lineNumber++
		for symbolIndex, c := range currentLine {
			char := string(c)
			if char == "." {
				continue
			}

			if !helpers.IsDigit(char) {
				checkTopCenter := true
				checkTopRight := true
				topLeftChar := string(previousLine[symbolIndex-1])
				if helpers.IsDigit(topLeftChar) {
					if helpers.IsDigit(string(previousLine[symbolIndex])) {
						checkTopCenter = false
						if helpers.IsDigit(string(previousLine[symbolIndex+1])) {
							checkTopRight = false
						}
					}

					numberStack := helpers.Stack{}
					StackLeft(previousLine, symbolIndex-2, &numberStack)
					i, _ := strconv.Atoi(string(previousLine[symbolIndex-1]))
					numberStack.Push(i)
					StackRight(previousLine, symbolIndex, &numberStack)

					total += numberStack.ConvertToNumber()
				}

				if checkTopCenter {
					topCenterChar := string(previousLine[symbolIndex])
					if helpers.IsDigit(topCenterChar) {
						checkTopRight = false

						numberStack := helpers.Stack{}
						StackRight(previousLine, symbolIndex, &numberStack)

						total += numberStack.ConvertToNumber()
					}
				}

				if checkTopRight {
					topRightChar := string(previousLine[symbolIndex+1])
					if helpers.IsDigit(topRightChar) {
						numberStack := helpers.Stack{}
						StackRight(previousLine, symbolIndex+1, &numberStack)

						total += numberStack.ConvertToNumber()
					}
				}

				leftChar := string(currentLine[symbolIndex-1])
				if helpers.IsDigit(leftChar) {
					numberStack := helpers.Stack{}
					StackLeft(currentLine, symbolIndex-1, &numberStack)

					total += numberStack.ConvertToNumber()
				}

				rightChar := string(currentLine[symbolIndex+1])
				if helpers.IsDigit(rightChar) {
					numberStack := helpers.Stack{}
					StackRight(currentLine, symbolIndex+1, &numberStack)

					total += numberStack.ConvertToNumber()
				}

				checkBottomCenter := true
				checkBottomRight := true
				bottomLeftChar := string(nextLine[symbolIndex-1])
				if helpers.IsDigit(bottomLeftChar) {
					if helpers.IsDigit(string(nextLine[symbolIndex])) {
						checkBottomCenter = false
						if helpers.IsDigit(string(nextLine[symbolIndex+1])) {
							checkBottomRight = false
						}
					}

					numberStack := helpers.Stack{}
					StackLeft(nextLine, symbolIndex-2, &numberStack)
					i, _ := strconv.Atoi(string(nextLine[symbolIndex-1]))
					numberStack.Push(i)
					StackRight(nextLine, symbolIndex, &numberStack)

					total += numberStack.ConvertToNumber()
				}

				if checkBottomCenter {
					bottomCenterChar := string(nextLine[symbolIndex])
					if helpers.IsDigit(bottomCenterChar) {
						checkBottomRight = false

						numberStack := helpers.Stack{}
						StackRight(nextLine, symbolIndex, &numberStack)

						total += numberStack.ConvertToNumber()
					}
				}

				if checkBottomRight {
					bottomRightChar := string(nextLine[symbolIndex+1])
					if helpers.IsDigit(bottomRightChar) {
						numberStack := helpers.Stack{}
						StackRight(nextLine, symbolIndex+1, &numberStack)

						total += numberStack.ConvertToNumber()
					}
				}
			}
		}

		previousLine = currentLine
		currentLine = nextLine
		scanner.Scan()
		nextLine = scanner.Text()
		fmt.Println(total)
	}
}

func StackLeft(text string, index int, s *helpers.Stack) {
	if index < 0 {

		return
	}
	if helpers.IsDigit(string(text[index])) {
		StackLeft(text, index-1, s)
		i, _ := strconv.Atoi(string(text[index]))
		s.Push(i)
	}
}

func StackRight(text string, index int, s *helpers.Stack) {
	if index >= 140 {
		return
	}
	if helpers.IsDigit(string(text[index])) {
		i, _ := strconv.Atoi(string(text[index]))
		s.Push(i)
		StackRight(text, index+1, s)
	}
}
