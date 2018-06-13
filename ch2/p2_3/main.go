package main

import (
	"fmt"
	"time"

	"github.com/samzhou1221/studygo/ch2/p2_3/popcount"
)

func main() {

	start := time.Now()

	popcount.PopCount(123456789)

	end := time.Since(start)

	fmt.Printf("PopCount: %v\n", end)

	start = time.Now()

	popcount.PopCount2(123456789)

	end = time.Since(start)

	fmt.Printf("PopCount2: %v\n", end)
}
