package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func RunDay17() {
	file, err := os.Open("./inputs/day17inputsmall.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	heatToEndMap := make(map[string]int)
	theMap := make(map[int]string)
	lineNumber := 0
	for scanner.Scan() {
		theMap[lineNumber] = scanner.Text()
		lineNumber++
	}

	FindShortestPath(&theMap, &heatToEndMap)

}

func getHeatHash(row int, col int) string {
	return strconv.Itoa(row*10) + strconv.Itoa(col)
}

func getHeatValue(row int, col int, theMap *map[int]string) int {
	if row >= len(*theMap) {
		return 0
	}

	if col >= len(*theMap) {
		return 0
	}

	return int((*theMap)[row][col]) - 48
}

func FindShortestPath(theMap *map[int]string, heatToEndMap *map[string]int) {

	for row := len(*theMap) - 1; row >= 0; row-- {
		// set destination heat
		heatHash := getHeatHash(row, len(*theMap)-1)
		heatLoss := getHeatValue(row, len(*theMap)-1, theMap)
		(*heatToEndMap)[heatHash] = heatLoss

		// set bottom row heat
		GoingRight(row, len(*theMap)-1, theMap, heatToEndMap)
	}
}

func GoingRight(row int, col int, theMap *map[int]string, heatToEndMap *map[string]int) {
	for i := col - 1; i >= 0; i-- {
		myHeatHash := getHeatHash(row, i)
		myHeatLoss := getHeatValue(row, i, theMap)

		oneRightHeatHash := getHeatHash(row, i+1)
		twoRightHeatHash := getHeatHash(row, i+2)
		threeRightHeatHash := getHeatHash(row, i+3)

		if threeRightHeatLoss, threeRightHashExists := (*heatToEndMap)[threeRightHeatHash]; threeRightHashExists {
			if existingHeatLoss, myHashExists := (*heatToEndMap)[myHeatHash]; myHashExists {
				newHeatLoss := threeRightHeatLoss + myHeatLoss
				if newHeatLoss < existingHeatLoss {
					(*heatToEndMap)[myHeatHash] = newHeatLoss
				}
			} else {
				(*heatToEndMap)[myHeatHash] = threeRightHeatLoss + myHeatLoss
			}
		}

		if twoRightHeatLoss, twoRightHashExists := (*heatToEndMap)[twoRightHeatHash]; twoRightHashExists {
			if existingHeatLoss, myHashExists := (*heatToEndMap)[myHeatHash]; myHashExists {
				newHeatLoss := twoRightHeatLoss + myHeatLoss
				if newHeatLoss < existingHeatLoss {
					(*heatToEndMap)[myHeatHash] = newHeatLoss
				}
			} else {
				(*heatToEndMap)[myHeatHash] = twoRightHeatLoss + myHeatLoss
			}
		}

		if oneRightHeatLoss, oneRightHashExists := (*heatToEndMap)[oneRightHeatHash]; oneRightHashExists {
			if existingHeatLoss, myHashExists := (*heatToEndMap)[myHeatHash]; myHashExists {
				newHeatLoss := oneRightHeatLoss + myHeatLoss
				if newHeatLoss < existingHeatLoss {
					(*heatToEndMap)[myHeatHash] = newHeatLoss
				}
			} else {
				(*heatToEndMap)[myHeatHash] = oneRightHeatLoss + myHeatLoss
			}
		}
	}
}
