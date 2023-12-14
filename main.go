package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	RunDay14()
	fmt.Println("Done in:", time.Now().Sub(start))
}
