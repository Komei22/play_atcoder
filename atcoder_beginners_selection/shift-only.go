package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	var a []int

	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := 0; i < n; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		a = append(a, num)
	}

	count := 0
	for {
		for i, val := range a {
			if val%2 != 0 {
				fmt.Println(count)
				return
			}
			a[i] = val / 2
		}
		count++
	}
}
