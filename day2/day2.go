package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RunDay2Part2() {
	file, err := os.Open("./day2/day2input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	powersSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()

		allPulls := lineText[strings.Index(lineText, ":")+2 : len(lineText)]
		pulls := strings.Split(allPulls, "; ")

		redMax := 0
		greenMax := 0
		blueMax := 0

		for _, pull := range pulls {
			redRegex, _ := regexp.Compile("([0-9]+) red")
			redBalls := redRegex.FindStringSubmatch(pull)
			if len(redBalls) > 1 {
				numRed, _ := strconv.Atoi(redBalls[1])
				redMax = max(redMax, numRed)
			}

			greenRegex, _ := regexp.Compile("([0-9]+) green")
			greenBalls := greenRegex.FindStringSubmatch(pull)
			if len(greenBalls) > 1 {
				numGreen, _ := strconv.Atoi(greenBalls[1])
				greenMax = max(greenMax, numGreen)
			}

			blueRegex, _ := regexp.Compile("([0-9]+) blue")
			blueBalls := blueRegex.FindStringSubmatch(pull)
			if len(blueBalls) > 1 {
				numBlue, _ := strconv.Atoi(blueBalls[1])
				blueMax = max(blueMax, numBlue)
			}
		}
		powersSum += redMax * greenMax * blueMax
	}
	fmt.Println(powersSum)
}

func RunDay2Part1() {
	file, err := os.Open("./day2/day2input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	redLimit := 12
	greenLimit := 13
	blueLimit := 14
	gameIdSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		lineRegex, _ := regexp.Compile("^Game ([0-9]+)")
		gameId := lineRegex.FindStringSubmatch(lineText)[1]
		gameNum, _ := strconv.Atoi(gameId)

		allPulls := lineText[strings.Index(lineText, ":")+2 : len(lineText)]
		pulls := strings.Split(allPulls, "; ")
		isGameValid := true
		for _, pull := range pulls {
			redRegex, _ := regexp.Compile("([0-9]+) red")
			redBalls := redRegex.FindStringSubmatch(pull)
			if len(redBalls) > 1 {
				numRed, _ := strconv.Atoi(redBalls[1])
				if numRed > redLimit {
					isGameValid = false
					break
				}
			}

			greenRegex, _ := regexp.Compile("([0-9]+) green")
			greenBalls := greenRegex.FindStringSubmatch(pull)
			if len(greenBalls) > 1 {
				numGreen, _ := strconv.Atoi(greenBalls[1])
				if numGreen > greenLimit {
					isGameValid = false
					break
				}
			}

			blueRegex, _ := regexp.Compile("([0-9]+) blue")
			blueBalls := blueRegex.FindStringSubmatch(pull)
			if len(blueBalls) > 1 {
				numBlue, _ := strconv.Atoi(blueBalls[1])
				if numBlue > blueLimit {
					isGameValid = false
					break
				}
			}
		}
		if isGameValid {
			gameIdSum += gameNum
		}

	}
	fmt.Println(gameIdSum)
}
