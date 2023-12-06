package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type MapNode struct {
	sourceStart      int
	destinationStart int
	rangeLength      int
	nextNode         *MapNode
}

func RunDay5Part2() {
	file, err := os.Open("./day5/day5input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seedLine := scanner.Text()[7:]
	seedsStrings := strings.Split(seedLine, " ")
	seeds := []int{}
	for _, v := range seedsStrings {
		seedInt, _ := strconv.Atoi(v)
		seeds = append(seeds, seedInt)
	}

	scanner.Scan()
	scanner.Scan()
	seedToSoilValues := ParseValues(scanner)

	scanner.Scan()
	soilToFertilizerValues := ParseValues(scanner)

	scanner.Scan()
	fertilizerToWaterValues := ParseValues(scanner)

	scanner.Scan()
	waterToLightValues := ParseValues(scanner)

	scanner.Scan()
	lightToTemperatureValues := ParseValues(scanner)

	scanner.Scan()
	temperatureToHumidityValues := ParseValues(scanner)

	scanner.Scan()
	humidityToLocationValues := ParseValues(scanner)

	seedToSoilMap := MakeMap(seedToSoilValues)
	soilToFertilizerMap := MakeMap(soilToFertilizerValues)
	fertilizationToWaterMap := MakeMap(fertilizerToWaterValues)
	waterToLightMap := MakeMap(waterToLightValues)
	lightToTemperatureMap := MakeMap(lightToTemperatureValues)
	temperatureToHumidityMap := MakeMap(temperatureToHumidityValues)
	humidityToLocationMap := MakeMap(humidityToLocationValues)

	minLocation := math.MaxInt
	for i := 0; i < len(seeds); i = i + 2 {
		seed := seeds[i]
		seedRange := seeds[i+1]

		for j := 0; j < seedRange; j++ {
			trySeed := seed + j
			soil := GetMapDestination(trySeed, seedToSoilMap)
			fertilizer := GetMapDestination(soil, soilToFertilizerMap)
			water := GetMapDestination(fertilizer, fertilizationToWaterMap)
			light := GetMapDestination(water, waterToLightMap)
			temperature := GetMapDestination(light, lightToTemperatureMap)
			humidity := GetMapDestination(temperature, temperatureToHumidityMap)
			location := GetMapDestination(humidity, humidityToLocationMap)

			minLocation = min(minLocation, location)
		}
	}
	fmt.Println(minLocation)
}

func RunDay5Part1() {
	file, err := os.Open("./day5/day5inputsmall.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seedLine := scanner.Text()[7:]
	seedsStrings := strings.Split(seedLine, " ")
	seeds := []int{}
	for _, v := range seedsStrings {
		seedInt, _ := strconv.Atoi(v)
		seeds = append(seeds, seedInt)
	}

	scanner.Scan()
	scanner.Scan()
	seedToSoilValues := ParseValues(scanner)

	scanner.Scan()
	soilToFertilizerValues := ParseValues(scanner)

	scanner.Scan()
	fertilizerToWaterValues := ParseValues(scanner)

	scanner.Scan()
	waterToLightValues := ParseValues(scanner)

	scanner.Scan()
	lightToTemperatureValues := ParseValues(scanner)

	scanner.Scan()
	temperatureToHumidityValues := ParseValues(scanner)

	scanner.Scan()
	humidityToLocationValues := ParseValues(scanner)

	seedToSoilMap := MakeMap(seedToSoilValues)
	soilToFertilizerMap := MakeMap(soilToFertilizerValues)
	fertilizationToWaterMap := MakeMap(fertilizerToWaterValues)
	waterToLightMap := MakeMap(waterToLightValues)
	lightToTemperatureMap := MakeMap(lightToTemperatureValues)
	temperatureToHumidityMap := MakeMap(temperatureToHumidityValues)
	humidityToLocationMap := MakeMap(humidityToLocationValues)

	minLocation := math.MaxInt
	for _, seed := range seeds {
		soil := GetMapDestination(seed, seedToSoilMap)
		fertilizer := GetMapDestination(soil, soilToFertilizerMap)
		water := GetMapDestination(fertilizer, fertilizationToWaterMap)
		light := GetMapDestination(water, waterToLightMap)
		temperature := GetMapDestination(light, lightToTemperatureMap)
		humidity := GetMapDestination(temperature, temperatureToHumidityMap)
		location := GetMapDestination(humidity, humidityToLocationMap)

		fmt.Println("Seed", seed, "Soil", soil)
		minLocation = min(minLocation, location)
	}
	fmt.Println(minLocation)
}

func GetMapDestination(sourceValue int, head *MapNode) int {
	currentNode := head
	for currentNode != nil {
		if sourceValue >= currentNode.sourceStart &&
			sourceValue < currentNode.sourceStart+currentNode.rangeLength {
			offset := sourceValue - currentNode.sourceStart
			return currentNode.destinationStart + offset
		} else {
			currentNode = currentNode.nextNode
		}
	}
	return sourceValue
}

func MakeMap(values [][]string) *MapNode {
	var returnMap *MapNode = nil
	lastNode := &MapNode{}

	for _, v := range values {
		destinationStart, _ := strconv.Atoi(v[0])
		sourceStart, _ := strconv.Atoi(v[1])
		rangeLength, _ := strconv.Atoi(v[2])

		newNode := MapNode{
			destinationStart: destinationStart,
			sourceStart:      sourceStart,
			rangeLength:      rangeLength,
		}
		if returnMap == nil {
			returnMap = &newNode
			lastNode = returnMap
		} else {
			lastNode.nextNode = &newNode
			lastNode = &newNode
		}
	}

	return returnMap
}

func ParseValues(scanner *bufio.Scanner) [][]string {
	parsedValues := [][]string{}
	scanner.Scan()
	textToParse := scanner.Text()
	for textToParse != "" {
		values := strings.Split(textToParse, " ")
		parsedValues = append(parsedValues, values)
		scanner.Scan()
		textToParse = scanner.Text()
	}

	return parsedValues
}
