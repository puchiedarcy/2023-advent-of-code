package main

import (
	"2023-advent-of-code/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunDay8Part1() {
	file, err := os.Open("./inputs/day8input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	directions := scanner.Text()

	scanner.Scan()
	nodes := make(map[string][]string)

	allNodes := []string{}
	for scanner.Scan() {
		lineText := scanner.Text()

		currentNode := lineText[0:3]
		leftNode := lineText[7:10]
		rightNode := lineText[12:15]

		nodes[currentNode] = []string{currentNode, leftNode, rightNode}
		if currentNode[2] == 'A' {
			allNodes = append(allNodes, currentNode)
		}
	}

	steps := 0
	stepsToZ := make(map[string]int)
	for len(stepsToZ) != len(allNodes) {
		for _, direction := range directions {
			steps++
			for i, node := range allNodes {
				if _, ok := stepsToZ[node]; ok {
					continue
				}

				if direction == 'L' {
					allNodes[i] = nodes[node][1]
				} else {
					allNodes[i] = nodes[node][2]
				}

				if allNodes[i][2] == 'Z' {
					stepsToZ[allNodes[i]] = steps
				}
			}
		}
	}

	stepNumbers := []int{}
	for _, v := range stepsToZ {
		stepNumbers = append(stepNumbers, v)
	}

	fmt.Println(helpers.FindLCM(stepNumbers))
}
