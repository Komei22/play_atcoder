package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

func sumOfDigit(n int) int {
	sum := 0
	for {
		digit := n % 10
		sum += digit
		n /= 10
		if n <= 0 {
			return sum
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N, A, B := nextInt(), nextInt(), nextInt()

	sum := 0
	for n := 1; n <= N; n++ {
		digitSum := sumOfDigit(n)
		if digitSum >= A && digitSum <= B {
			sum += n
		}
	}
	fmt.Println(sum)
}
