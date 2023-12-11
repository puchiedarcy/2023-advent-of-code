package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	RunDay11()
	fmt.Println("Done in:", time.Now().Sub(start))
}
