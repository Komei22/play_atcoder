package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	cStr := a + b
	c, _ := strconv.ParseFloat(cStr, 64)
	cSqrt := math.Sqrt(c)
	if cSqrt-math.Trunc(cSqrt) == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
