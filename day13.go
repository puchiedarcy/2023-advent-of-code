package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunDay13() {
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
			mirrorIndex, found := findSymmetry(rowMap)
			if !found {
				mirrorIndex, found = findSymmetry(colMap)
				colTotal += mirrorIndex + 1
			} else {
				rowTotal += mirrorIndex + 1
			}

			fmt.Println("found", mirrorIndex+1)
			row = 0
			rowMap = make(map[int]string)
			colMap = make(map[int]string)
			continue
		}

		rowMap[row] = lineText
		row++
		for k, v := range lineText {
			colMap[k] += string(v)
		}
	}

	fmt.Println(rowMap)
	fmt.Println(colMap)
	fmt.Println("search LAST rows")
	mirrorIndex, found := findSymmetry(rowMap)
	if !found {
		fmt.Println("search cols")
		mirrorIndex, found = findSymmetry(colMap)
		colTotal += mirrorIndex + 1
	} else {
		rowTotal += mirrorIndex + 1
	}
	fmt.Println("found", mirrorIndex+1)

	theTotal := rowTotal*100 + colTotal
	fmt.Println(theTotal)
}

func findSymmetry(theMap map[int]string) (int, bool) {
	mirrorIndex := 0
	for mirrorIndex < len(theMap)-1 {
		if theMap[mirrorIndex] == theMap[mirrorIndex+1] {

			leftIndex := mirrorIndex - 1
			rightIndex := mirrorIndex + 2

			isCompleteMirror := true
			for leftIndex >= 0 && rightIndex < len(theMap) {
				if theMap[leftIndex] == theMap[rightIndex] {
					leftIndex--
					rightIndex++
					continue
				} else {
					isCompleteMirror = false
					break
				}
			}

			if isCompleteMirror {
				return mirrorIndex, true
			}
		}
		mirrorIndex++
	}
	return 0, false
}
