package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	RunDay16()
	fmt.Println("Done in:", time.Now().Sub(start))
}
