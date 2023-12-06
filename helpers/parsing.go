package helpers

import (
	"regexp"
	"strconv"
)

func ParseInts(textLine string) []int {
	intRegex, _ := regexp.Compile("-?\\d+")
	ints := intRegex.FindAllStringSubmatch(textLine, -1)
	returnVal := []int{}
	for _, v := range ints {
		theInt, _ := strconv.Atoi(v[0])
		returnVal = append(returnVal, theInt)
	}
	return returnVal
}
