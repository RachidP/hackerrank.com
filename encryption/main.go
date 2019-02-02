package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	in, _ := reader.ReadString('\n')
	//fmt.Println("input = ", in)
	in = strings.Replace(in, " ", "", -1)
	copyIn := in
	fmt.Println("input = ", in)
	sizeIn := float64(len(in))
	//rows := math.Floor(math.Sqrt(sizeIn))
	columns := math.Ceil(math.Sqrt(sizeIn))
	//fmt.Printf("rows = %v,  columns = %v\n", rows, columns)
	//rowsInt := int(rows)
	columnsInt := int(columns)
	//a := make([]string, rowsInt)
	//fmt.Println("size a= ", len(a))
	//i := 0
	//for i < rowsInt-1 {
	//	a[i] = in[:columnsInt]
	//	in = in[columnsInt:]
	//	i++
	//}
	//a[i] = in

	tmp := []byte(copyIn)
	//fmt.Println(a, in, copyIn, tmp)
	for i := 0; i < columnsInt; i++ {
		for x := i; x < len(tmp); {
			fmt.Printf("%s", string(tmp[x]))
			x += columnsInt
		}
		fmt.Printf(" ")
	}

}
