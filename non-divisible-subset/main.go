<<<<<<< HEAD
<<<<<<< HEAD
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n, k uint64

	fmt.Scanf("%v %v\n", &n, &k)
	reader := bufio.NewReader(os.Stdin)
	in, _ := reader.ReadString('\n')
	strArr := strings.Fields(in)

	fmt.Println(n, k)
	fmt.Println(strArr)
	remaindersS := make([]uint64, k)
	for _, v := range strArr {

		numero, _ := strconv.ParseUint(v, 10, 64)
		//conto quanti resti sn presenti nell'intervallo 0...k-1
		remaindersS[numero%k]++
	}
	fmt.Println(remaindersS)

	//c := math.Min(float64(remaindersS[0]), 1)
	c := myMin(remaindersS[0], 1)
	for i := uint64(1); i < (k/2)+1; i++ {
		if i != k-i {
			//c += math.Max(float64(remaindersS[i]), float64(remaindersS[k-i]))
			c += myMax(remaindersS[i], remaindersS[k-i])

		}

	}
	if k%2 == 0 {
		c++
	}
	fmt.Println(c)

}

func myMax(x1, x2 uint64) uint64 {
	if x1 >= x2 {
		return x1
	}
	return x2

}
func myMin(x1, x2 uint64) uint64 {
	if x1 <= x2 {
		return x1
	}
	return x2
}
=======
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n, k uint64

	fmt.Scanf("%v %v\n", &n, &k)
	reader := bufio.NewReader(os.Stdin)
	in, _ := reader.ReadString('\n')
	strArr := strings.Fields(in)

	fmt.Println(n, k)
	fmt.Println(strArr)
	remaindersS := make([]uint64, k)
	for _, v := range strArr {

		numero, _ := strconv.ParseUint(v, 10, 64)
		//conto quanti resti sn presenti nell'intervallo 0...k-1
		remaindersS[numero%k]++
	}
	fmt.Println(remaindersS)

	//c := math.Min(float64(remaindersS[0]), 1)
	c := myMin(remaindersS[0], 1)
	for i := uint64(1); i < (k/2)+1; i++ {
		if i != k-i {
			//c += math.Max(float64(remaindersS[i]), float64(remaindersS[k-i]))
			c += myMax(remaindersS[i], remaindersS[k-i])

		}

	}
	if k%2 == 0 {
		c++
	}
	fmt.Println(c)

}

func myMax(x1, x2 uint64) uint64 {
	if x1 >= x2 {
		return x1
	}
	return x2

}
func myMin(x1, x2 uint64) uint64 {
	if x1 <= x2 {
		return x1
	}
	return x2
}
<<<<<<< HEAD
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
