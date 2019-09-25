package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

func main() {
	A, B, C, X := nextLine(), nextLine(), nextLine(), nextLine()
	count := 0
	for a := 0; a <= A; a++ {
		for b := 0; b <= B; b++ {
			for c := 0; c <= C; c++ {
				if (a*500 + b*100 + c*50) == X {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
