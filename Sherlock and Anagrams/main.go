<<<<<<< HEAD
<<<<<<< HEAD
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

//https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	data := []rune(s)
	size := len(data)
	var m map[string]int32
	var res int32

	for i := 1; i < size; i++ {
		//make a map to salve the string and the number of repetition of that string
		m = make(map[string]int32)

		//slice the string in substrings
		for j := 0; j < size-i+1; j++ {

			// calcolate and get the new substring
			subString := []rune(s[j : j+i])

			//order the substring
			sort.Slice(subString, func(i, j int) bool {
				return subString[i] < subString[j]
			})

			//add the substring in the map and update the value
			m[string(subString)]++

		}
		//read the map
		for _, v := range m {
			//get the number of occurency of the substring
			if v > 1 {
				//calcolate Triangular number (wiki)
				res += (v - 1) * v / 2
			}
		}

	}

	//fmt.Println(res)

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
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

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
=======
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
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

//https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	data := []rune(s)
	size := len(data)
	var m map[string]int32
	var res int32

	for i := 1; i < size; i++ {
		//make a map to salve the string and the number of repetition of that string
		m = make(map[string]int32)

		//slice the string in substrings
		for j := 0; j < size-i+1; j++ {

			// calcolate and get the new substring
			subString := []rune(s[j : j+i])

			//order the substring
			sort.Slice(subString, func(i, j int) bool {
				return subString[i] < subString[j]
			})

			//add the substring in the map and update the value
			m[string(subString)]++

		}
		//read the map
		for _, v := range m {
			//get the number of occurency of the substring
			if v > 1 {
				//calcolate Triangular number (wiki)
				res += (v - 1) * v / 2
			}
		}

	}

	//fmt.Println(res)

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
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

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
<<<<<<< HEAD
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
