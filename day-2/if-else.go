package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	if 7%2 == 0 {
		fmt.Println("7 is even!")
	} else {
		fmt.Println("7 is odd!")
	}
	elapsed := time.Since(start)
	fmt.Println(" \n", elapsed)
	fmt.Println(" Calculating times with for loops till 10 :::: ")
	startTen := time.Now()
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	elapsedTen := time.Since(startTen)
	fmt.Println(" \n", elapsedTen)
	fmt.Println(" Calculating times with for loops 1000 :::: ")
	start1000 := time.Now()
	sum1000 := 0
	for i := 0; i < 1000; i++ {
		sum1000 += i
	}
	elapsed1000 := time.Since(start1000)
	fmt.Println(" \n", elapsed1000)
}