package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

func next(sc *bufio.Scanner) int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	// bufio.ScanWordsはスペース区切りだけではなく改行でも区切ってくれる。便利！
	// 改行区切りのみのときはデフォルトのままでよし。
	// ただし、トークンが長いときは読めないので、readlineを使おう。
	sc.Split(bufio.ScanWords)
	N := next(sc)

	var a []int
	for i := 0; i < N; i++ {
		a = append(a, next(sc))
	}

	var alice, bob int
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			alice += a[i]
		} else {
			bob += a[i]
		}
	}
	fmt.Println(alice-bob)
}
