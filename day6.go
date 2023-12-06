package main

import (
	"2023-advent-of-code/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunDay6Part1() {
	file, err := os.Open("./inputs/day6input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	times := helpers.ParseInts(scanner.Text())
	scanner.Scan()
	distances := helpers.ParseInts(scanner.Text())

	allWaysToWin := 1
	for i := range times {
		middleTime := times[i] / 2
		firstWinningTime := FindFirstWinningTime(middleTime, 0, times[i], distances[i])
		waysToWin := (middleTime - firstWinningTime + 1) * 2
		if times[i]%2 == 0 {
			waysToWin--
		}
		allWaysToWin *= waysToWin
		fmt.Println("---")
		fmt.Println(waysToWin)
		fmt.Println("---")
	}
	fmt.Println(allWaysToWin)
}

func FindFirstWinningTime(holdTime int, leftHoldTime int, totalTime int, winningDistance int) int {
	runTime := totalTime - holdTime
	distance := holdTime * runTime

	fmt.Println("Left", leftHoldTime, "Right", holdTime)
	if leftHoldTime > holdTime {
		fmt.Println("invalid test")
		return -1
	}

	if distance > winningDistance {
		fmt.Println("Go Deeper")
		if holdTime == leftHoldTime {
			fmt.Println("Found Same", holdTime)
			return holdTime
		}

		firstWinningTime := FindFirstWinningTime((holdTime+leftHoldTime)/2, leftHoldTime, totalTime, winningDistance)
		if firstWinningTime == -1 {
			firstWinningTime = FindFirstWinningTime((holdTime+(holdTime/2))/2, holdTime/2+1, totalTime, winningDistance)
			if firstWinningTime != -1 {
				return firstWinningTime
			} else {
				firstWinningTime = FindFirstWinningTime(holdTime-1, (holdTime+(holdTime/2))/2+1, totalTime, winningDistance)
				if firstWinningTime != -1 {
					return firstWinningTime
				} else {
					fmt.Println("Found", holdTime)
					return holdTime
				}
			}
		} else {
			return firstWinningTime
		}
	}
	fmt.Println("Fail")
	return -1
}
