package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	oClock  = " o' clock"
	to      = " to "
	past    = " past "
	minute  = " minute"
	minutes = " minutes"
)

var numbers = []string{"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"quarter",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
	"twenty",
	"twenty one",
	"twenty two",
	"twenty three",
	"twenty four",
	"twenty five",
	"twenty six",
	"twenty seven",
	"twenty eight",
	"twenty nine",
	"half",
}

// Complete the timeInWords function below.
func timeInWords(h int32, m int32) string {

	if m <= 30 {
		switch m {
		case 0:
			return numbers[h] + oClock
		case 1:
			return numbers[m] + minute + past + numbers[h]
		case 15:
			return numbers[m] + past + numbers[h]
		case 30:
			return numbers[m] + past + numbers[h]
		default:
			return numbers[m] + minutes + past + numbers[h]

		}

	}

	switch diff := 60 - m; diff {

	case 15:
		return numbers[diff] + to + numbers[h+1]

	default:
		return numbers[diff] + minutes + to + numbers[h+1]

	}

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

	hTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	h := int32(hTemp)

	mTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := timeInWords(h, m)

	fmt.Fprintf(writer, "%s\n", result)

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
