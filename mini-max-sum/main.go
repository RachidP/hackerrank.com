package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const MaxUint32 = 1<<32 - 1
	var sizeV = 5
	var vet [5]uint32
	reader := bufio.NewReader(os.Stdin)
	in, _ := reader.ReadString('\n')
	strArr := strings.Fields(in)
	for index := 0; index < sizeV; index++ {
		tmp, _ := strconv.ParseUint(strArr[index], 10, 64)
		vet[index] = uint32(tmp)
	}
	fmt.Println(vet)
	var tmp uint32
	var min uint32
	min = MaxUint32
	var max uint32
	for i := 0; i < sizeV; i++ {
		for j := 0; j < sizeV; j++ {
			if i == j {
				continue
			}
			tmp += vet[j]
		}
		if min > tmp {
			min = tmp
		}
		if max < tmp {
			max = tmp
		}
		tmp = 0
	}
	fmt.Println(min, max)

}
