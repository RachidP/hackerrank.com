package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var m (map[int64]int64)

// Complete the stoneDivision function below.
func stoneDivision(n int64, s []int64) int64 {
	// fmt.Println(n, s)
	m = make(map[int64]int64)
	var res int64
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	res = calc(n, s)
	return res

}

func calc(n int64, s []int64) int64 {
	var res int64
	if v, found := m[n]; found {
		return v
	}

	for _, v := range s {
		if v == n {
			break
		}

		if r := n % v; r == 0 {
			nSubProblems := n / v
			tmp := (calc(v, s) * nSubProblems) + 1
			if tmp > res {
				res = tmp
			}
		}
	}
	m[n] = res
	return res
}
func main() {

	f, err := os.Open("input.txt")
	checkError(err)
	defer f.Close()
	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout, err := os.Create("OUTPUT_PATH")
	checkError(err)
	defer stdout.Close()
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nm := strings.Split(readLine(reader), " ")

		n, err := strconv.ParseInt(nm[0], 10, 64)
		checkError(err)

		mTemp, err := strconv.ParseInt(nm[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		sTemp := strings.Split(readLine(reader), " ")

		var s []int64

		for i := 0; i < int(m); i++ {
			sItem, err := strconv.ParseInt(sTemp[i], 10, 64)
			checkError(err)
			s = append(s, sItem)
		}

		result := stoneDivision(n, s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
