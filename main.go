package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	RunDay5Part2()
	fmt.Println("Done in: ", time.Now().Sub(start))
}
