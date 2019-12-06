package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func next(sc *bufio.Scanner) int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	N := next(sc)

	var d []int
	for i := 0; i < N; i++ {
		d = append(d, next(sc))
	}

	var dan []int
	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	for i, val := range d {
		if i == 0 {
			dan = append(dan, val)
			continue
		}
		if val < dan[len(dan)-1] {
			dan = append(dan, val)
		}
	}
	fmt.Println(len(dan))
}
