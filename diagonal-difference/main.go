<<<<<<< HEAD
<<<<<<< HEAD
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	var sizeA int
	var in string
	var diagA, diagB int
	fmt.Scanf("%v\n", &sizeA)

	data := make([][]int, sizeA)

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < sizeA; i++ {
		in, _ = reader.ReadString('\n')
		strArr := strings.Fields(in)
		for y := 0; y < sizeA; y++ {
			tmp, _ := strconv.Atoi(strArr[y])
			data[i] = append(data[i], tmp)

		}
	}
	fmt.Printf("%#v \n", data)

	for i, y := 0, sizeA-1; i < sizeA; i++ {
		diagA += data[i][i]
		diagB += data[i][y]
		/*	fmt.Printf("i=%v   y= %v \n", i, y)
			fmt.Printf("diagA: %v += data[%v][%v]=%v\n", diagA, i, i, data[i][i])
			fmt.Printf("diagB: %v += data[%v][%v]=%v\n", diagB, i, y, data[i][y])
		*/
		y--

	}

	res := float64(diagA - diagB)
	//	fmt.Println(res)
	res = math.Abs(res)
	fmt.Println(res)
}
=======
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	var sizeA int
	var in string
	var diagA, diagB int
	fmt.Scanf("%v\n", &sizeA)

	data := make([][]int, sizeA)

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < sizeA; i++ {
		in, _ = reader.ReadString('\n')
		strArr := strings.Fields(in)
		for y := 0; y < sizeA; y++ {
			tmp, _ := strconv.Atoi(strArr[y])
			data[i] = append(data[i], tmp)

		}
	}
	fmt.Printf("%#v \n", data)

	for i, y := 0, sizeA-1; i < sizeA; i++ {
		diagA += data[i][i]
		diagB += data[i][y]
		/*	fmt.Printf("i=%v   y= %v \n", i, y)
			fmt.Printf("diagA: %v += data[%v][%v]=%v\n", diagA, i, i, data[i][i])
			fmt.Printf("diagB: %v += data[%v][%v]=%v\n", diagB, i, y, data[i][y])
		*/
		y--

	}

	res := float64(diagA - diagB)
	//	fmt.Println(res)
	res = math.Abs(res)
	fmt.Println(res)
}
<<<<<<< HEAD
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
