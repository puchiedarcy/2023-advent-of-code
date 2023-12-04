package day4

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func RunDay4Part1() {
	file, err := os.Open("./day4/day4input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()

		winningNumbers := []int{}
		numTemp, _ := strconv.Atoi(strings.TrimSpace(lineText[10:12]))
		winningNumbers = append(winningNumbers, numTemp)
		numTemp, _ = strconv.Atoi(strings.TrimSpace(lineText[13:15]))
		winningNumbers = append(winningNumbers, numTemp)
		numTemp, _ = strconv.Atoi(strings.TrimSpace(lineText[16:18]))
		winningNumbers = append(winningNumbers, numTemp)
		numTemp, _ = strconv.Atoi(strings.TrimSpace(lineText[19:21]))
		winningNumbers = append(winningNumbers, numTemp)
		numTemp, _ = strconv.Atoi(strings.TrimSpace(lineText[22:24]))
		winningNumbers = append(winningNumbers, numTemp)
		numTemp, _ = strconv.Atoi(strings.TrimSpace(lineText[25:27]))
		winningNumbers = append(winningNumbers, numTemp)
		numTemp, _ = strconv.Atoi(strings.TrimSpace(lineText[28:30]))
		winningNumbers = append(winningNumbers, numTemp)
		numTemp, _ = strconv.Atoi(strings.TrimSpace(lineText[31:33]))
		winningNumbers = append(winningNumbers, numTemp)
		numTemp, _ = strconv.Atoi(strings.TrimSpace(lineText[34:36]))
		winningNumbers = append(winningNumbers, numTemp)
		numTemp, _ = strconv.Atoi(strings.TrimSpace(lineText[37:39]))
		winningNumbers = append(winningNumbers, numTemp)

		matchingNumbers := 0
		for i := 0; i < 25; i++ {
			numberGiven, _ := strconv.Atoi(strings.TrimSpace(lineText[42+i*3 : 44+i*3]))
			for _, v := range winningNumbers {
				if numberGiven == v {
					matchingNumbers++
					break
				}
			}
		}
		if matchingNumbers != 0 {
			sum += int(math.Pow(float64(2), float64(matchingNumbers-1)))
		}
	}
	fmt.Println(sum)
}
