package main

//https://www.hackerrank.com/challenges/crush/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Solution with index calcolation: very fast
// when a query to add V is given in range [a, b] we will add V to arr[a] and –V to arr[b+1] now if we want
// to get the actual values of array we will convert the above array into prefix sum array
func arrayManipulation(n int32, queries [][]int32) int64 {
	//	fmt.Println(queries)

	var max, tmp, k int64
	var i, startIndex, endIndex int32
	arr := make([]int64, n+1)

	for _, row := range queries {
		startIndex, endIndex, k = row[0], row[1], int64(row[2])

		arr[startIndex] += k
		if endIndex != n {
			arr[endIndex+1] -= k
		}
	}

	//	fmt.Printf("%v %v %v \n", startIndex, endIndex, k)
	for i = 1; i <= n; i++ {
		tmp += arr[i]
		if tmp > max {
			max = tmp
		}

	}
	fmt.Println(max)

	return max
}

//Simle solution but take  long time
func arrayManipulationSimpleSoution(n int32, queries [][]int32) int64 {
	fmt.Println(queries)

	arr := make([]int64, n+1)

	var max, tmp int64
	for _, row := range queries {
		startIndex, endIndex, k := row[0], row[1], int64(row[2])

		//	fmt.Printf("%v %v %v \n", startIndex, endIndex, k)
		for i := startIndex; i <= endIndex; i++ {
			arr[endIndex] += k
			arr[i] += k
			tmp = arr[i]
			if tmp > max {
				max = tmp
			}
			tmp = arr[endIndex]
			if tmp > max {
				max = tmp
			}
			endIndex--
		}

	}
	fmt.Println(max)

	return max
}

// Complete the arrayManipulation function below.
func arrayManipulation2(n int32, queries [][]int32) int64 {
	arr := make([]int64, n+1)

	var max, tmp int64
	for _, row := range queries {
		startIndex, endIndex, k := row[0], row[1], int64(row[2])

		//  fmt.Printf("%v %v %v \n", startIndex, endIndex, k)
		for i := startIndex; i <= endIndex; i++ {

			arr[i] += k
			tmp = arr[i]
			if tmp > max {
				max = tmp
			}
		}

	}
	//  fmt.Println(max)

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

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

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
