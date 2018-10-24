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

var m map[int64]int64

// Complete the decibinaryNumbers function below.
func decibinaryNumbers(x int64) int64 {

	// for i := int64(0); i < 15; i++ {
	// 	var tot int64
	// 	input := getEachIndividualDigit(i)

	// 	for i, v := range input {

	// 		tot += v * m[int64(i)]
	// 	}
	// 	fmt.Println("total=", tot)
	// 	fmt.Println("*************************")

	// }

	// indice := make([]int64, 10)
	// values := make([]int64, 10)
	// for x := int64(0); x < 1100000; x++ {
	// 	for i := int64(0); i < 10; i++ {
	// 		var tot int64
	// 		input := getEachIndividualDigit(i + (x * 10))

	// 		for i, v := range input {

	// 			tot += v * m[int64(i)]
	// 		}

	// 		indice[i] = i + (x * 10)
	// 		values[i] = tot

	// 	}
	// 	fmt.Println(indice)
	// 	fmt.Println(values)
	// 	fmt.Println("-----")

	// }

	// for i := int64(0); i < 15; i++ {
	// 	var tot int64
	// 	input := getEachIndividualDigit(i)

	// 	for i, v := range input {

	// 		tot += v * m[int64(i)]
	// 	}
	// 	fmt.Println("total=", tot)
	// 	fmt.Println("*************************")

	// }
	// fmt.Println("***********************************")
	// input := getEachIndividualDigit(110)
	// var tot int64
	// for i, v := range input {

	// 	tot += v * m[int64(i)]
	// }
	// fmt.Println("total=", tot)
	//_ = getDecimal(0)
	for i := int64(0); i < 20; i++ {
		fmt.Printf("\n %v : %v : %v ", i, i, getDecimal(i))
	}
	x = 10999969
	fmt.Printf("\ndecimal of %v = %v \n", x, getDecimal(x))
	return -1

}

func getDecimal(decibinary int64) int64 {
	if decibinary < 10 {
		return decibinary
	}
	var tot int64
	i := int(math.Log10(float64(decibinary))) //i=log(num)
	iPow10 := int64(math.Pow10(i))
	quotient := decibinary / iPow10 //q= num/(10^i)
	rest := decibinary % iPow10     //q= num/(10^i)
	iPow2 := int64(math.Pow(2, float64(i)))
	if rest > 9 {
		tot = getDecimal(rest) + (iPow2 * quotient)
		return tot
	}

	tot = m[rest] + (iPow2 * quotient)

	return tot
}
func getDecimal2(decibinary int64) int64 {
	if decibinary < 10 {
		return decibinary
	}
	i := int(math.Log10(float64(decibinary))) //i=log(num)
	iPow10 := int64(math.Pow10(i))
	quotient := decibinary / iPow10 //q= num/(10^i)
	rest := decibinary % iPow10     //q= num/(10^i)
	//	tot := m[rest] + (2 * quotient)
	// num := float64(decibinary)
	// i := int(math.Log10(num))                 //i=log(num)
	// quotient := int64(num / math.Pow10(i))    //q= num/(10^i)
	// rest := int64(num) % int64(math.Pow10(i)) //q= num/(10^i)
	// tot := m[rest] + (2 * quotient)
	iPow2 := int64(math.Pow(2, float64(i)))
	tot := m[rest] + (iPow2 * quotient)

	return tot
}

func getEachIndividualDigit(z int64) []int64 {
	res := make([]int64, 0, 100)
	var tot int64
	var i float64
	// tmp := z ^ (z & 9)
	//z>0

	for z != 0 {
		//resto
		resto := z % 10

		tot += resto * int64(math.Pow(2, i))
		//iteger part
		z /= 10
		//salve result
		res = append(res, resto)

	}

	//	fmt.Println(res)
	return res

}
func fillMap() {
	// PowersOf10 := []int64{1, 10, 100, 1000, 10000, 100000,
	// 	1000000, 10000000, 100000000, 1000000000}
	// _ = PowersOf10
	// m = make(map[int64]int64, 100)
	// for i := uint(0); i < 20; i++ {
	// 	m[int64(i)] = 1 << i
	// }
	m = map[int64]int64{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}

	fmt.Println(m)
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

	fillMap()

	for qItr := 0; qItr < int(q); qItr++ {
		x, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		result := decibinaryNumbers(x)

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


func subsetLexRank(n int64, t []int64) int64{
	

	r :=0
	for i=1;i<=n;i++{
		if i
	}
}