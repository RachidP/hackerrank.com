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

	var sizeA int
	var in string
	//var data []int
	fmt.Scanf("%v\n", &sizeA)
	//fmt.Scanf("%v", &vett)
	//
	data := make([]int, sizeA)
	/*scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in = scanner.Text()
		break

	}*/
	reader := bufio.NewReader(os.Stdin)

	in, _ = reader.ReadString('\n')
	fmt.Println(sizeA)
	fmt.Println(in)
	strArr := strings.Fields(in)
	var sum int
	for i, v := range strArr {
		data[i], _ = strconv.Atoi(v)
		sum += data[i]
		//fmt.Println(i, v)
	}
	fmt.Println(data)
	fmt.Println(sum)
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

	var sizeA int
	var in string
	//var data []int
	fmt.Scanf("%v\n", &sizeA)
	//fmt.Scanf("%v", &vett)
	//
	data := make([]int, sizeA)
	/*scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in = scanner.Text()
		break

	}*/
	reader := bufio.NewReader(os.Stdin)

	in, _ = reader.ReadString('\n')
	fmt.Println(sizeA)
	fmt.Println(in)
	strArr := strings.Fields(in)
	var sum int
	for i, v := range strArr {
		data[i], _ = strconv.Atoi(v)
		sum += data[i]
		//fmt.Println(i, v)
	}
	fmt.Println(data)
	fmt.Println(sum)
}
<<<<<<< HEAD
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
