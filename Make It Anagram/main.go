package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var seen map[int32]int32

// Complete the stepPerms function below.
func stepPerms(n int32) int32 {
	var res int32
	if n < 0 {

		return 0
	}
	if v, ok := seen[n]; ok {
		return v
	}
	res = stepPerms(n-1) + stepPerms(n-2) + stepPerms(n-3)
	seen[n] = res

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

	sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	s := int32(sTemp)

	for sItr := 0; sItr < int(s); sItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)
		seen = make(map[int32]int32)
		seen[0] = 1

		res := stepPerms(n)

		fmt.Fprintf(writer, "%d\n", res)
		//	fmt.Println("finished: look at the file \"OUTPUT_PATH\"")
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
