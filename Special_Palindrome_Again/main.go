package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the substrCount function below.
func substrCount(n int32, s string) int64 {
	//	fmt.Println(n, s)
	var res, count int64
	seq := make([]int64, 0, n)
	cur := rune(s[0])
	count++

	//.All of the characters are the same, e.g. aaa.

	//make a list of tuples that can form a palindromic substring
	//and add to res the combination of these substring (eq. aaa)
	for _, v := range s[1:] {
		if v == cur {
			count++
			continue
		}

		seq = append(seq, count)
		res += (count * (count + 1)) / 2

		cur = v
		count = 1

	}
	seq = append(seq, count)
	res += (count * (count + 1)) / 2

	//All characters except the middle one are the same, e.g. aadaa.
	sizeS := len(seq)
	iPrev := seq[0] - 1 //index of previous letter
	for i := 1; i < sizeS-1; i++ {
		iNext := iPrev + seq[i] + seq[i+1]
		if seq[i] == 1 && s[iPrev] == s[iNext] {
			res += int64(math.Min(float64(seq[i-1]), float64(seq[i+1])))
		}
		iPrev += seq[i]
	}

	//fmt.Println(len(seq), seq, res)

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

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	result := substrCount(n, s)

	fmt.Fprintf(writer, "%d\n", result)

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
