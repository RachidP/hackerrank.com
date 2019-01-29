package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the maxSubsetSum function below.
func maxSubsetSum(arr []int32) int32 {

	sizeA := int32(len(arr))
	m := map[int32]int32{-2: 0, -1: 0, 0: arr[0]}
	max := arr[0]

	for i := int32(1); i < sizeA; i++ {

		tmpMax := arr[i]

		relAss := arr[i] + m[i-2]
		if tmpMax < relAss {
			tmpMax = relAss
		}

		if tmpMax < m[i-2] {
			tmpMax = m[i-2]
		}

		m[i] = tmpMax

		if max < tmpMax {
			max = tmpMax
		} else {
			m[i] = max
		}

	}
	fmt.Println("max:", max)
	return max
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

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := maxSubsetSum(arr)

	fmt.Fprintf(writer, "%d\n", res)

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
