package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	RunDay15()
	fmt.Println("Done in:", time.Now().Sub(start))
}
