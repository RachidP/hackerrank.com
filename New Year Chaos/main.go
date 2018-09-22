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

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	//sizQ := int32(len(q))
	var tot, temp int32
	diff := float64(0)
	fmt.Println(q)
	for i, v := range q {
		temp++
		diff = float64(v - temp)
		if diff > 2 {

			fmt.Println("Too chaotic")
			return
		}

		for j := int(math.Max(0, float64(v-2))); j < i; j++ {
			if q[j] > q[i] {
				tot++
			}

		}

	}
	//fmt.Println(oneMove, twoMove, oneMove-twoMove)
	fmt.Println(tot)

}

// Complete the minimumBribes function below.
func minimumBribes3(q []int32) {
	var tot, temp int32
	diff := float64(0)
	// fmt.Println(q)
	for _, v := range q {
		temp++
		if v > temp {
			diff = math.Abs(float64(v - temp))
			if diff > 2 {
				fmt.Println("Too chaotic")
				return
			}
			tot += int32(diff)

		}

	}
	//fmt.Println(oneMove, twoMove, oneMove-twoMove)
	fmt.Println(tot)

}
func minimumBribes2(q []int32) {
	//sizQ := int32(len(q))
	var temp, oneMove, twoMove int32
	diff := float64(0)
	fmt.Println(q)
	for _, v := range q {
		temp++
		if v <= temp+2 && v >= temp-2 {

			diff = math.Abs(float64(v - temp))
			if diff == 1 {
				oneMove++
			} else {
				twoMove++
			}

		} else {
			fmt.Println("Too chaotic")
			return

		}

	}
	//fmt.Println(oneMove, twoMove, oneMove-twoMove)
	fmt.Println(oneMove - twoMove)

}
func main() {
	f, err := os.Open("input.txt")
	checkError(err)
	defer f.Close()
	reader := bufio.NewReaderSize(f, 1024*1024)
	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
