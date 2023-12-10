package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	RunDay9Part1()
	fmt.Println("Done in:", time.Now().Sub(start))
}
