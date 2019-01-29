<<<<<<< HEAD
package main

//https://www.hackerrank.com/challenges/count-triplets-1/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {

	mB := make(map[int64]int64)
	mC := make(map[int64]int64)
	var tot int64

	fmt.Println(arr)

	for _, v := range arr {

		if val, ok := mC[v]; ok {
			tot += val
		}

		if val, ok := mB[v]; ok {
			mC[v*r] += val
		}
		mB[v*r]++

	}
	fmt.Println(tot)
	return tot
}
func countTriplets2(arr []int64, r int64) int64 {

	sizeArr := len(arr)
	m := make(map[int64]int64)
	var x, y, z, tot int64
	var ok bool

	fmt.Println(arr)

	for _, v := range arr {
		m[v]++
	}

	for i := int64(sizeArr - 1); i >= 0; i -= z {

		tmp := arr[i]
		z, _ = m[tmp]          //get the number of repetition of the element on the map
		firstResult := tmp / r //calcolate the first Result
		if y, ok = m[firstResult]; !ok {
			//i -= z
			continue
		}

		secondResult := firstResult / r
		if x, ok = m[secondResult]; !ok {

			continue
		}
		// if firstResult == secondResult {
		// 	tot += z
		// } else {

		tot += (x * y * z)
		//	}

	}
	fmt.Println(tot)

	return tot
}

func main() {

	f, err := os.Open("input.txt")
	checkError(err)
	defer f.Close()
	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout, err := os.Create("OUTPUT_PATH")

	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
=======
package main

//https://www.hackerrank.com/challenges/count-triplets-1/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {

	mB := make(map[int64]int64)
	mC := make(map[int64]int64)
	var tot int64

	fmt.Println(arr)

	for _, v := range arr {

		if val, ok := mC[v]; ok {
			tot += val
		}

		if val, ok := mB[v]; ok {
			mC[v*r] += val
		}
		mB[v*r]++

	}
	fmt.Println(tot)
	return tot
}
func countTriplets2(arr []int64, r int64) int64 {

	sizeArr := len(arr)
	m := make(map[int64]int64)
	var x, y, z, tot int64
	var ok bool

	fmt.Println(arr)

	for _, v := range arr {
		m[v]++
	}

	for i := int64(sizeArr - 1); i >= 0; i -= z {

		tmp := arr[i]
		z, _ = m[tmp]          //get the number of repetition of the element on the map
		firstResult := tmp / r //calcolate the first Result
		if y, ok = m[firstResult]; !ok {
			//i -= z
			continue
		}

		secondResult := firstResult / r
		if x, ok = m[secondResult]; !ok {

			continue
		}
		// if firstResult == secondResult {
		// 	tot += z
		// } else {

		tot += (x * y * z)
		//	}

	}
	fmt.Println(tot)

	return tot
}

func main() {

	f, err := os.Open("input.txt")
	checkError(err)
	defer f.Close()
	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout, err := os.Create("OUTPUT_PATH")

	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
