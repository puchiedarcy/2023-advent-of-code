package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
				s := strings.Replace(string(v), "#", "1", -1)
				colMap[k] += strings.Replace(s, ".", "0", -1)
			}
		}
	}

	r, c := doThePuzzle(rowMap, colMap)
	rowTotal += r
	colTotal += c

	fmt.Println(rowTotal*100 + colTotal)
}

func doThePuzzle(rowMap map[int]string, colMap map[int]string) (int, int) {
	oldRowMirror, _ := findExistingSymmetry(rowMap)
	fmt.Println("Check Rows, but skip", oldRowMirror)
	newMirror, newFound := findNewSymmetry(rowMap, oldRowMirror)

	if newFound {
		return newMirror, 0
	}

	oldColMirror, _ := findExistingSymmetry(colMap)
	fmt.Println("Check Cols, but skip", oldColMirror)
	newColMirror, newColFound := findNewSymmetry(colMap, oldColMirror)
	if !newColFound {
		log.Fatalln(rowMap)
	}
	return 0, newColMirror
}

func findNewSymmetry(bitMap map[int]string, skip int) (int, bool) {
	hammings := make(map[int]int)
	zeroHams := []int{}
	oneHams := []int{}
	for i := 0; i < len(bitMap)-1; i++ {
		h := hamming([]byte(bitMap[i]), []byte(bitMap[i+1]))
		hammings[i+1] = h
		if h == 0 {
			zeroHams = append(zeroHams, i+1)
		} else if h == 1 {
			oneHams = append(oneHams, i+1)
		}
	}

	fmt.Println(hammings)
	fmt.Println(zeroHams)

	newMirror, found := checkHams(zeroHams, hammings, bitMap, skip, false)
	if found {
		return newMirror, found
	}

	fmt.Println(oneHams)
	newMirror, found = checkHams(oneHams, hammings, bitMap, skip, true)
	return newMirror, found
}

func checkHams(mirrorList []int, hammings map[int]int, bitMap map[int]string, skip int, usedSmudge bool) (int, bool) {
	for i := 0; i < len(mirrorList); i++ {
		if mirrorList[i] == skip {
			continue
		}
		usedSmudge := usedSmudge
		left := mirrorList[i] - 1
		right := mirrorList[i] + 1

		isMirror := true
		for left > 0 || right <= len(bitMap) {

			fmt.Println(left, mirrorList[i], right)
			leftEdge := false
			if left <= 0 {
				leftEdge = true
			}

			rightEdge := false
			if right > len(bitMap) {
				rightEdge = true
			}

			hD := 0
			if !leftEdge && !rightEdge {
				hD = int(math.Abs(float64(hammings[left] - hammings[right])))
			} else if leftEdge {
				hD = hammings[right]
			} else if rightEdge {
				hD = hammings[left]
			} else {
				isMirror = false
				break
			}

			if hD == 0 {
				left--
				right++
				continue
			} else if hD == 1 && !usedSmudge {
				usedSmudge = true
				left--
				right++
				continue
			} else {
				isMirror = false
				break
			}
		}

		if isMirror {
			fmt.Println("FOUND", mirrorList[i])
			return mirrorList[i], true
		}
	}
	return 0, false
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
