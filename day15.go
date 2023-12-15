package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	focal int
	label string
}

type LensBox struct {
	lenses  []Lens
	lensLoc map[string]int
}

func RunDay15() {
	file, err := os.Open("./inputs/day15input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	currentValue := 0
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lineText := scanner.Text()
	steps := strings.Split(lineText, ",")

	allBoxes := []LensBox{}
	for i := 0; i < 256; i++ {
		lensBox := LensBox{lensLoc: make(map[string]int)}
		allBoxes = append(allBoxes, lensBox)
	}

	for _, step := range steps {
		for _, c := range step {
			if c == '=' || c == '-' {
				break
			}
			currentValue += int(c)
			currentValue *= 17
			currentValue %= 256
		}

		boxNumber := currentValue
		isDash := step[len(step)-1] == '-'
		if !isDash {
			labelFocal := strings.Split(step, "=")
			label := labelFocal[0]
			focal, _ := strconv.Atoi(labelFocal[1])

			lens := Lens{label: label, focal: focal}

			indexToReplace, exists := allBoxes[boxNumber].lensLoc[label]
			if exists {
				allBoxes[boxNumber].lenses[indexToReplace].focal = focal
			} else {
				allBoxes[boxNumber].lenses = append(allBoxes[boxNumber].lenses, lens)
				allBoxes[boxNumber].lensLoc[label] = len(allBoxes[boxNumber].lenses) - 1
			}
		} else {
			label := step[:len(step)-1]
			indexToRemove, exists := allBoxes[boxNumber].lensLoc[label]
			if exists {
				tmpList := allBoxes[boxNumber].lenses[:indexToRemove]
				for _, v := range allBoxes[boxNumber].lenses[indexToRemove+1:] {
					allBoxes[boxNumber].lensLoc[v.label] -= 1
				}
				tmpList = append(tmpList, allBoxes[boxNumber].lenses[indexToRemove+1:]...)
				allBoxes[boxNumber].lenses = tmpList
				delete(allBoxes[boxNumber].lensLoc, label)

			}
		}
		currentValue = 0
	}

	focusingPower := 0
	for boxNumber := 0; boxNumber < len(allBoxes); boxNumber++ {
		for lensNumber := 0; lensNumber < len(allBoxes[boxNumber].lenses); lensNumber++ {
			focusingPower += (boxNumber + 1) * (lensNumber + 1) * allBoxes[boxNumber].lenses[lensNumber].focal
		}
	}
	fmt.Println(focusingPower)
}
