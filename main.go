package main

import (
	"2023-advent-of-code/day5"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	day5.RunDay5Part2()
	fmt.Println("Done in: ", time.Now().Sub(start))
}
