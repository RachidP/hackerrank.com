package main

import (
	"fmt"
	"strconv"
)

var memo = make(map[string]string)

func main() {

	var n string
	var k int
	fmt.Scanf("%s %v \n", &n, &k)

	//for i := uint64(0); i < k; i++ {
	//	tot += n
	//}

	fmt.Println(digitSum(n, k))

}

func digitSum(n string, k int) int {
	var sum int
	for _, v := range n {
		sum += int(v - '0')
		sum %= 9
	}
	sum *= k
	if sum == 0 {
		return 9
	}

	return sum % 9

}
func calcolate2(x string) string {
	dim := len(x)
	var tot string
	if dim <= 1 {
		return x
	}
	if dim == 2 {
		num1, _ := strconv.Atoi(x[:1])
		num2, _ := strconv.Atoi(x[1:])
		tmp := num2 + num1
		if tmp >= 10 {
			tot = strconv.Itoa(tmp)
			num1, _ = strconv.Atoi(tot[:1])
			num2, _ = strconv.Atoi(tot[1:])
			tmp = num2 + num1
		}

		tot = strconv.Itoa(tmp)
		return tot

	}

	dim = len(x)

	x = calcolate(x[:dim/2]) + calcolate(x[dim/2:])
	x = calcolate(x)

	return x
}

func calcolate(x string) string {
	dim := len(x)
	if v, ok := memo[x]; ok {
		return v
	}
	var tot string
	if dim <= 1 {
		return x
	}
	if dim == 2 {
		num1, _ := strconv.Atoi(x[:1])
		num2, _ := strconv.Atoi(x[1:])
		tmp := num2 + num1
		if tmp >= 10 {
			tot = strconv.Itoa(tmp)
			num1, _ = strconv.Atoi(tot[:1])
			num2, _ = strconv.Atoi(tot[1:])
			tmp = num2 + num1
		}

		tot = strconv.Itoa(tmp)
		return tot

	}

	dim = len(x)

	tmp1 := calcolate(x[:dim/2])
	memo[x[:dim/2]] = tmp1
	tmp2 := calcolate(x[dim/2:])
	memo[x[dim/2:]] = tmp2
	x = tmp1 + tmp2
	x = calcolate(x)

	return x
}
