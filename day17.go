package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type HeatNode struct {
	heatLoss      int
	up            *HeatNode
	upHeatLoss    int
	down          *HeatNode
	downHeatLoss  int
	left          *HeatNode
	leftHeatLoss  int
	right         *HeatNode
	rightHeatLoss int
}

func NewHeatNode(heatLoss int) *HeatNode {
	newHeatNode := HeatNode{heatLoss: heatLoss}
	return &newHeatNode
}

func RunDay17() {
	file, err := os.Open("./inputs/day17inputsmall.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	theMap := make(map[int]string)
	lineNumber := 0
	for scanner.Scan() {
		theMap[lineNumber] = scanner.Text()
		lineNumber++
	}

	destinationNode := NewHeatNode(getHeatValue(len(theMap)-1, len(theMap)-1, &theMap))

	currentNode := destinationNode
	row := len(theMap) - 1
	// set all left and right on row
	for col := len(theMap[row]) - 2; col >= 0; col-- {
		leftNode := NewHeatNode(getHeatValue(row, col, &theMap))
		currentNode.left = leftNode
		leftNode.right = currentNode
		currentNode = leftNode
	}

	// set all up down on col
	currentNode = destinationNode
	rowNode := currentNode
	for col := len(theMap) - 1; col >= 0; col-- {
		currentNode = rowNode
		for upRow := len(theMap) - 2; upRow >= 0; upRow-- {
			upNode := NewHeatNode(getHeatValue(upRow, col, &theMap))
			currentNode.up = upNode
			upNode.down = currentNode
			currentNode = upNode
		}
		rowNode = rowNode.left
	}

	// connect all rows left and right
	bottomRowNodes := []*HeatNode{}
	currentNode = destinationNode
	for currentNode != nil {
		bottomRowNodes = append(bottomRowNodes, currentNode)
		currentNode = currentNode.left
	}

	currentNode = bottomRowNodes[0]
	nextNode := bottomRowNodes[1]
	for nextNode != nil {
		rightNode := currentNode.up
		leftNode := nextNode.up

		for leftNode != nil {
			rightNode.left = leftNode
			leftNode.right = rightNode

			rightNode = rightNode.up
			leftNode = leftNode.up
		}

		toBeCurrentNode := nextNode
		nextNode = nextNode.left
		currentNode = toBeCurrentNode
	}

	currentNode = destinationNode
	for currentNode != nil {
		tmpNextCurrent := currentNode.left
		if tmpNextCurrent == nil {
			break
		} else {
			currentNode = tmpNextCurrent
		}
	}

	for currentNode != nil {
		tmpNextCurrent := currentNode.up
		if tmpNextCurrent == nil {
			break
		} else {
			currentNode = tmpNextCurrent
		}
	}

	sourceNode := currentNode

	fmt.Println(currentNode)
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
