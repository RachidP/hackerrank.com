package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var sizeA int

	var in string

	fmt.Scanf("%v\n", &sizeA)
	data := make([]int32, sizeA)
	reader := bufio.NewReader(os.Stdin)
	in, _ = reader.ReadString('\n')
	strArr := strings.Fields(in)

	for i, v := range strArr {
		tmp, _ := strconv.ParseInt(v, 10, 32)
		data[i] = int32(tmp)

	}
	fmt.Println(data)
	sort.Slice(data, func(i, j int) bool {
		return data[i] <= data[j]
	})

	i := int32(sizeA - 1)
	max := data[i]
	var tot int32

	for i >= 0 && data[i] == max {

		tot++
		i--
	}
	fmt.Println(i)
	fmt.Println(tot)
}
