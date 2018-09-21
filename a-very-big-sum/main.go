package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var sizeA int
	var in string
	fmt.Scanf("%v\n", &sizeA)

	reader := bufio.NewReader(os.Stdin)
	in, _ = reader.ReadString('\n')
	strArr := strings.Fields(in)
	var sum uint64
	for _, v := range strArr {
		tmp, _ := strconv.ParseUint(v, 10, 64)
		sum += tmp
		//fmt.Println(i, v)
	}

	fmt.Println(sum)
}
