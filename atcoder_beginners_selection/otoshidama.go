package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func next(sc *bufio.Scanner) int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	N, Y := next(sc), next(sc)

	impossible := true
	for i := 0; i <= N; i++ {
		for j := 0; j <= N-i; j++ {
			if (10000*i + 5000*j + 1000*(N-i-j)) == Y {
				impossible = false
				fmt.Printf("%d %d %d\n", i, j, N-i-j)
				break
			}
		}
		if !impossible {
			break
		}
	}
	if impossible {
		fmt.Println("-1 -1 -1")
	}
}
