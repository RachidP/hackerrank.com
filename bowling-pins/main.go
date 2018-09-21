package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nim = make([]int, 301)

func main() {

	var t, result int
	var in string
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d\n", &t)

	for t > 0 {

		_, _ = reader.ReadString('\n')
		in, _ = reader.ReadString('\n')
		in = strings.TrimSuffix(in, "\r\n")
		piles := inizializeSlice(in)

		result = 0
		for _, v := range piles {
			result ^= grundy(v)
		}
		printWinner(result)

		t--
	}

}

func printWinner(result int) {
	if result > 0 {
		fmt.Println("WIN")
	} else {
		fmt.Println("LOSE")
	}
}

func inizializeSlice(gameStr string) []int {
	s := make([]int, 1)
	pinsI := 'I'
	prevRune := []rune(gameStr)[0]
	var count int

	for _, v := range gameStr {
		if v == prevRune {

			if pinsI == v {

				s[count]++
			}

		} else {
			if pinsI == v {

				s[count]++
			} else {
				count++
				s = append(s, 0)
			}
		}
		prevRune = v
	}

	if s[len(s)-1] == 0 {
		return s[:len(s)-1]
	}
	return s
}

func calculateMex(m map[int]int) int {
	var mex int

	for _, ok := m[mex]; ok; _, ok = m[mex] {
		mex++
	}
	return mex
}

func grundy(n int) int {

	if nim[n] != 0 {
		return nim[n]
	}
	if n == 0 {
		nim[0] = 0
		return 0
	}
	if n == 1 {
		nim[1] = 1
		return 1
	}
	if n == 2 {
		nim[2] = 2
		return 2
	}
	//	nim[3] = 3

	h := make(map[int]int)
	for i, j, k := 1, 0, 0; i <= n; i++ {
		var x, y int
		if i <= n/2 {
			x = grundy(j)
			y = grundy(n - j - 2)
			j++
		} else {
			x = grundy(k)
			y = grundy(n - k - 1)
			k++
		}

		h[x^y] = x ^ y
	}

	nim[n] = calculateMex(h)

	return nim[n]

}
