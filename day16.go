package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func RunDay16() {
	file, err := os.Open("./inputs/day16input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	theMap := make(map[int]string)
	lineNumber := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		theMap[lineNumber] = lineText
		lineNumber++
	}

	energizedMap := make(map[int]string)
	for i := 0; i < len(theMap); i++ {
		energizedMap[i] = strings.Repeat(".", len(theMap[i]))
	}

	alreadyBeamed := make(map[int]bool)

	FollowBeam(0, 0, LEFT, &theMap, &energizedMap, &alreadyBeamed)

	energyLevel := 0
	for i := 0; i < len(energizedMap); i++ {
		energyLevel += strings.Count(energizedMap[i], "#")
	}
	fmt.Println(energyLevel)
}

func FollowBeam(row int, col int, enteringFrom int, theMap *map[int]string, energizedMap *map[int]string, alreadyBeamed *map[int]bool) {
	beamHash := row*10000 + col*10 + enteringFrom
	_, exists := (*alreadyBeamed)[beamHash]
	if exists {
		return
	}

	(*alreadyBeamed)[beamHash] = true

	if enteringFrom == LEFT {
		energizeCell(row, col, energizedMap)
		for col < len((*theMap)[row]) {
			if (*theMap)[row][col] == '.' ||
				(*theMap)[row][col] == '-' {
				energizeCell(row, col, energizedMap)
				col++
			} else {
				break
			}
		}

		if col < len((*theMap)[row]) {
			energizeCell(row, col, energizedMap)
			switch (*theMap)[row][col] {
			case '|':
				if row-1 >= 0 {
					FollowBeam(row-1, col, DOWN, theMap, energizedMap, alreadyBeamed)
				}
				if row+1 < len(*theMap) {
					FollowBeam(row+1, col, UP, theMap, energizedMap, alreadyBeamed)
				}
			case '/':
				if row-1 >= 0 {
					FollowBeam(row-1, col, DOWN, theMap, energizedMap, alreadyBeamed)
				}
			case '\\':
				if row+1 < len(*theMap) {
					FollowBeam(row+1, col, UP, theMap, energizedMap, alreadyBeamed)
				}
			}
		}
	} else if enteringFrom == RIGHT {
		energizeCell(row, col, energizedMap)
		for col >= 0 {
			if (*theMap)[row][col] == '.' ||
				(*theMap)[row][col] == '-' {
				energizeCell(row, col, energizedMap)
				col--
			} else {
				break
			}
		}

		if col >= 0 {
			energizeCell(row, col, energizedMap)
			switch (*theMap)[row][col] {
			case '|':
				if row-1 >= 0 {
					FollowBeam(row-1, col, DOWN, theMap, energizedMap, alreadyBeamed)
				}
				if row+1 < len(*theMap) {
					FollowBeam(row+1, col, UP, theMap, energizedMap, alreadyBeamed)
				}
			case '/':
				if row+1 < len(*theMap) {
					FollowBeam(row+1, col, UP, theMap, energizedMap, alreadyBeamed)
				}
			case '\\':
				if row-1 >= 0 {
					FollowBeam(row-1, col, DOWN, theMap, energizedMap, alreadyBeamed)
				}
			}
		}
	} else if enteringFrom == UP {
		energizeCell(row, col, energizedMap)
		for row < len(*theMap) {
			if (*theMap)[row][col] == '.' ||
				(*theMap)[row][col] == '|' {
				energizeCell(row, col, energizedMap)
				row++
			} else {
				break
			}
		}

		if row != len(*theMap) {
			energizeCell(row, col, energizedMap)
			switch (*theMap)[row][col] {
			case '-':
				if col-1 >= 0 {
					FollowBeam(row, col-1, RIGHT, theMap, energizedMap, alreadyBeamed)
				}
				if col+1 < len(*theMap) {
					FollowBeam(row, col+1, LEFT, theMap, energizedMap, alreadyBeamed)
				}
			case '/':
				if col-1 >= 0 {
					FollowBeam(row, col-1, RIGHT, theMap, energizedMap, alreadyBeamed)
				}
			case '\\':
				if col+1 < len(*theMap) {
					FollowBeam(row, col+1, LEFT, theMap, energizedMap, alreadyBeamed)
				}
			}
		}
	} else if enteringFrom == DOWN {
		energizeCell(row, col, energizedMap)
		for row >= 0 {
			if (*theMap)[row][col] == '.' ||
				(*theMap)[row][col] == '|' {
				energizeCell(row, col, energizedMap)
				row--
			} else {
				break
			}
		}

		if row >= 0 {
			energizeCell(row, col, energizedMap)
			switch (*theMap)[row][col] {
			case '-':
				if col-1 >= 0 {
					FollowBeam(row, col-1, RIGHT, theMap, energizedMap, alreadyBeamed)
				}
				if col+1 < len(*theMap) {
					FollowBeam(row, col+1, LEFT, theMap, energizedMap, alreadyBeamed)
				}
			case '/':
				if col+1 < len(*theMap) {
					FollowBeam(row, col+1, LEFT, theMap, energizedMap, alreadyBeamed)
				}
			case '\\':

				if col-1 >= 0 {
					FollowBeam(row, col-1, RIGHT, theMap, energizedMap, alreadyBeamed)
				}
			}
		}
	}
}

func energizeCell(row int, col int, energizedMap *map[int]string) {
	replacementRow := (*energizedMap)[row][:col] + "#"
	if col+1 < len((*energizedMap)[row]) {
		replacementRow += (*energizedMap)[row][col+1:]
	}
	(*energizedMap)[row] = replacementRow
}
