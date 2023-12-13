package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func RunDay13Part2() {
	file, err := os.Open("./inputs/day13input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	colTotal := 0
	rowTotal := 0
	row := 0
	rowMap := make(map[int]string)
	colMap := make(map[int]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		if lineText == "" {
			r, c := doThePuzzle(rowMap, colMap)
			rowTotal += r
			colTotal += c

			colMap = make(map[int]string)
			rowMap = make(map[int]string)
			row = 0
			continue
		} else {
			inBits := strings.Replace(lineText, "#", "1", -1)
			inBits = strings.Replace(inBits, ".", "0", -1)
			rowMap[row] = inBits
			row++
			for k, v := range lineText {
				colMap[k] += string(v)
			}
		}
	}

	r, c := doThePuzzle(rowMap, colMap)
	rowTotal += r
	colTotal += c

	fmt.Println(rowTotal*100 + colTotal)
}

func doThePuzzle(rowMap map[int]string, colMap map[int]string) (int, int) {
	mirror, found := findExistingSymmetry(rowMap)
	if found {
		return mirror, 0
	}

	mirror, found = findExistingSymmetry(colMap)
	return 0, mirror
}

func findExistingSymmetry(bitMap map[int]string) (int, bool) {
	for i := 0; i < len(bitMap)-1; i++ {
		distance := hamming([]byte(bitMap[i]), []byte(bitMap[i+1]))
		if distance == 0 {
			left := i
			right := i + 1
			isMirror := true
			for left >= 0 && right < len(bitMap) {
				if 0 == hamming([]byte(bitMap[left]), []byte(bitMap[right])) {
					left--
					right++
				} else {
					isMirror = false
					break
				}
			}

			if isMirror {
				return i + 1, true
			}
		} else {
			continue
		}
	}
	return 0, false
}

func hamming(a, b []byte) int {
	diff := 0
	for i := 0; i < len(a); i++ {
		b1 := a[i]
		b2 := b[i]
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if (b1 & mask) != (b2 & mask) {
				diff++
			}
		}
	}
	return diff
}
