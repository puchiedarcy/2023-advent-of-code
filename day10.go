package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	UNKNOWN = iota
	LEFT
	RIGHT
	UP
	DOWN
)

func RunDay10() {
	file, err := os.Open("./inputs/day10input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pipeMap := []string{}
	startRow := 0
	startCol := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		pipeMap = append(pipeMap, lineText)
		start := strings.IndexByte(lineText, 'S')
		if start != -1 {
			startCol = start
		}
		if startCol == -1 {
			startRow++
		}
	}

	dir1, dir2 := findStartDirections(&pipeMap, startRow, startCol)

	distance := TakeNextStepInBothDirections(
		&pipeMap, 0,
		startRow, startCol, dir1,
		startRow, startCol, dir2)
	fmt.Println(distance)
}

func TakeNextStepInBothDirections(
	theMap *[]string, steps int,
	pos1Row int, pos1Col int, dir1 int,
	pos2Row int, pos2Col int, dir2 int) int {

	if pos1Row == pos2Row &&
		pos1Col == pos2Col &&
		steps > 0 {
		return steps
	}

	next1Row, next1Col := calculateNextPos(pos1Row, pos1Col, dir1)
	next1Dir := calculateNextDir((*theMap)[next1Row][next1Col], dir1)

	next2Row, next2Col := calculateNextPos(pos2Row, pos2Col, dir2)
	next2Dir := calculateNextDir((*theMap)[next2Row][next2Col], dir2)

	return TakeNextStepInBothDirections(
		theMap, steps+1,
		next1Row, next1Col, next1Dir,
		next2Row, next2Col, next2Dir)
}

func calculateNextPos(row int, col int, dir int) (int, int) {
	nextRow := row
	nextCol := col

	switch dir {
	case LEFT:
		nextCol--
	case RIGHT:
		nextCol++
	case UP:
		nextRow--
	case DOWN:
		nextRow++
	}

	return nextRow, nextCol
}

func calculateNextDir(c byte, prevDir int) int {
	if c == '-' {
		if prevDir == RIGHT {
			return RIGHT
		}
		return LEFT
	}

	if c == '7' {
		if prevDir == RIGHT {
			return DOWN
		}
		return LEFT
	}

	if c == '|' {
		if prevDir == UP {
			return UP
		}
		return DOWN
	}

	if c == 'J' {
		if prevDir == DOWN {
			return LEFT
		}
		return UP
	}

	if c == 'L' {
		if prevDir == DOWN {
			return RIGHT
		}
		return UP
	}

	if c == 'F' {
		if prevDir == UP {
			return RIGHT
		}
		return DOWN
	}

	return UNKNOWN
}

func findStartDirections(theMap *[]string, row int, col int) (int, int) {
	left := col - 1
	right := col + 1
	up := row - 1
	down := row + 1

	dir1 := UNKNOWN
	dir2 := UNKNOWN
	if left > 0 {
		if (*theMap)[row][left] != '.' {
			dir1 = LEFT
		}
	}

	if right < len(*theMap) {
		if (*theMap)[row][right] != '.' {
			if dir1 == UNKNOWN {
				dir1 = RIGHT
			} else {
				dir2 = RIGHT
			}
		}
	}

	if up > 0 {
		if (*theMap)[up][col] != '.' {
			if dir1 == UNKNOWN {
				dir1 = UP
			} else {
				dir2 = UP
			}
		}
	}

	if down < len(*theMap) {
		if (*theMap)[down][col] != '.' {
			dir2 = DOWN
		}
	}

	return dir1, dir2
}
