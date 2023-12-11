package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type Galaxy struct {
	xCoord int
	yCoord int
}

func RunDay11() {
	file, err := os.Open("./inputs/day11input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	galaxies := []Galaxy{}
	emptyRowsForExpansion := []int{}
	emptyColsForExpansion := []int{}
	colsWithAGalaxy := make(map[int]bool)

	scanner := bufio.NewScanner(file)
	currentRow := 0
	for scanner.Scan() {
		foundGalaxyOnThisRow := false
		for i, j := range scanner.Text() {
			if j == '#' {
				foundGalaxyOnThisRow = true
				galaxies = append(galaxies, Galaxy{xCoord: currentRow, yCoord: i})
				colsWithAGalaxy[i] = true
			}
		}

		if !foundGalaxyOnThisRow {
			emptyRowsForExpansion = append(emptyRowsForExpansion, currentRow)

		}
		currentRow++
	}

	for c := 0; c < currentRow; c++ {
		_, exists := colsWithAGalaxy[c]
		if !exists {
			emptyColsForExpansion = append(emptyColsForExpansion, c)
		}
	}

	sort.Ints(emptyRowsForExpansion)
	sort.Ints(emptyColsForExpansion)

	totalDistance := 0

	for i := 0; i < len(galaxies); i++ {
		galaxy1 := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			galaxy2 := galaxies[j]

			leftX := min(galaxy1.xCoord, galaxy2.xCoord)
			rightX := max(galaxy1.xCoord, galaxy2.xCoord)

			emptyRowsCrossed := 0
			for r := 0; r < len(emptyRowsForExpansion); r++ {
				emptyRow := emptyRowsForExpansion[r]
				if leftX < emptyRow && emptyRow < rightX {
					emptyRowsCrossed += 999999
				}
			}

			upperY := min(galaxy1.yCoord, galaxy2.yCoord)
			lowerY := max(galaxy1.yCoord, galaxy2.yCoord)

			emptyColsCrossed := 0
			for c := 0; c < len(emptyColsForExpansion); c++ {
				emptyCol := emptyColsForExpansion[c]
				if upperY < emptyCol && emptyCol < lowerY {
					emptyColsCrossed += 999999
				}
			}

			distance :=
				int(math.Abs(float64(galaxy2.xCoord-galaxy1.xCoord))) +
					int(math.Abs(float64(galaxy2.yCoord-galaxy1.yCoord))) +
					emptyRowsCrossed + emptyColsCrossed
			totalDistance += distance
		}
	}

	fmt.Println(totalDistance)
}
