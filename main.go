package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	RunDay8Part1()
	fmt.Println("Done in:", time.Now().Sub(start))
}
