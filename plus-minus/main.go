package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var sizeA float64
	//var tot float32
	var in string
	fmt.Scanf("%v\n", &sizeA)

	reader := bufio.NewReader(os.Stdin)
	in, _ = reader.ReadString('\n')
	strArr := strings.Fields(in)

	var pos, neg, zeros float64

	for _, v := range strArr {
		tmp, _ := strconv.ParseFloat(v, 64)
		if tmp > 0 {

			pos++

		} else if tmp < 0 {

			neg++
		} else {
			zeros++
		}

	}

	fmt.Println(pos / sizeA)
	fmt.Println(neg / sizeA)
	fmt.Println(zeros, zeros/sizeA)
	//	tmp := float32(pos) / float32(sizeA)
	//tmp := float32(pos / sizeA)
	//fmt.Printf("tmp= %v\n", tmp)
}
