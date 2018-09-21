package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var memo = make(map[string]bool)

func main() {
	var t int
	var pass, loginAttemp string
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d\n", &t)
	for t > 0 {
		//	fmt.Scanln(&n)
		_, _ = reader.ReadString('\n')
		//tmp = strings.TrimSuffix(tmp, "\r\n")
		//n, _ := strconv.Atoi(tmp)
		//fmt.Scanf("%d\n", &n)
		pass, _ = reader.ReadString('\n')
		//pass = strings.TrimSuffix(pass, "\r\n")
		passwords := strings.Fields(pass)
		loginAttemp, _ = reader.ReadString('\n')
		loginAttemp = strings.TrimSuffix(loginAttemp, "\r\n")
		//	fmt.Printf("passwords: %#v \n", passwords)
		//	fmt.Printf("loginAttempt: %#v \n", loginAttemp)
		//resolve(passwords, loginAttemp)
		////resolve2(passwords, loginAttemp)
		///*	find := resolveLast(passwords, loginAttemp, 0)
		//	if !find {
		//		fmt.Println("WRONG PASSWORD")
		//	}*/
		//
		crack(passwords, loginAttemp)
		for k := range memo {
			delete(memo, k)
		}

		t--
	}
	//passwords := []string{"ab", "abcd", "cd"}
	//loginAttemp := "abcd"

}

func resolve(passwords []string, loginAttemp string) {

	tmpString := loginAttemp
	for _, v := range passwords {
		size := len(v)
		for {
			if indice := strings.Index(tmpString, v); indice != -1 {
				canc := []byte(tmpString)
				tmpString = string(append(canc[:indice], tmpString[indice+size:]...))
			} else {
				break
			}
		}
		loginAttemp = strings.Replace(loginAttemp, v, v+" ", -1)

	}
	if tmpString == "" {
		fmt.Printf(loginAttemp)
	} else {
		fmt.Printf("WRONG PASSWORD")
	}
	fmt.Printf("\n")

}

func resolve2(passwords []string, loginAttemp string) {

	tmpString := loginAttemp
	soluz := make([]int, len(passwords))
	for i := len(passwords) - 1; i >= 0; i-- {
		parola := passwords[i]
		size := len(passwords[i])
		for {
			if indice := strings.Index(tmpString, parola); indice != -1 {
				canc := []byte(tmpString)
				tmpString = string(append(canc[:indice], tmpString[indice+size:]...))
				soluz[i]++

			} else {
				break
			}
		}
		//fmt.Println("parola= ", parola)

	}
	if tmpString == "" {
		for i, v := range passwords {
			if soluz[i] != 0 {
				loginAttemp = strings.Replace(loginAttemp, v, v+" ", soluz[i])
			}

		}
		fmt.Printf(loginAttemp)
	} else {
		fmt.Printf("WRONG PASSWORD")
	}
	fmt.Printf("\n")

}

func resolveLast(passwords []string, loginAttemp string, k int) bool {

	var find bool

	if k == len(passwords) {

		return rachid(passwords, loginAttemp)
	}
	for i := k; i < len(passwords); i++ {
		passwords[k], passwords[i] = passwords[i], passwords[k]
		find = resolveLast(passwords, loginAttemp, k+1)
		if find {
			break
		}
		passwords[k], passwords[i] = passwords[i], passwords[k]
	}

	return find
}

func rachid(passwords []string, loginAttemp string) bool {
	var solution string
	tmpString := loginAttemp
	for _, v := range passwords {
		size := len(v)
		for {
			if indice := strings.Index(tmpString, v); indice != -1 {
				solution = solution + v + " "
				canc := []byte(tmpString)
				tmpString = string(append(canc[:indice], tmpString[indice+size:]...))
			} else {
				break
			}
		}

	}

	if tmpString == "" {
		fmt.Println(solution)
		return true
	}
	return false

}
func resolveLast2(passwords []string, loginAttemp string, chosen []string, size int) bool {
	var find bool

	if size == 0 {
		return rachid(chosen, loginAttemp)
	}

	sizePasswords := len(passwords)
	for i := 0; i < sizePasswords; i++ {
		chosen = append(chosen, passwords[i])
		find = resolveLast2(passwords, loginAttemp, chosen, size-1)
		chosen = chosen[:len(chosen)-1]

		if find {
			return find
		}
	}
	return find
}

func crack(passwords []string, loginAttemp string) {
	var res []string

	_crack(passwords, loginAttemp, &res)
	if len(res) > 0 {
		for _, v := range res {
			fmt.Printf("%v ", v)

		}
		fmt.Printf("\n")
	} else {
		fmt.Println("WRONG PASSWORD")
	}
}
func _crack(keys []string, password string, res *[]string) bool {
	if len(password) == 0 {
		return true
	}
	if _, ok := memo[password]; ok {
		return false
	}

	for _, key := range keys {

		if len(key) <= len(password) && password[:len(key)] == key {
			*res = append(*res, key)
			memo[password] = true

			if _crack(keys, password[len(key):], res) {
				return true
			}
			slice := *res
			//res = *res[:len(*res)-1]
			*res = slice[:len(slice)-1]

		}
	}
	return false

}

func prova() {
	var vet []int

	fmt.Println(vet)

	calcola(&vet)
	fmt.Println(vet)

}
func calcola(vet *[]int) {

	*vet = append(*vet, 10, 20)
}
