<<<<<<< HEAD
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
		_, _ = reader.ReadString('\n')
		pass, _ = reader.ReadString('\n')
		passwords := strings.Fields(pass)
		loginAttemp, _ = reader.ReadString('\n')
		loginAttemp = strings.TrimSuffix(loginAttemp, "\r\n")
		crack(passwords, loginAttemp)
		for k := range memo {
			delete(memo, k)
		}
		t--
	}

	//passwords := []string{"ab", "abcd", "cd"}
	//loginAttemp := "abcd"

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

			*res = slice[:len(slice)-1]

		}
	}
	return false

}
=======
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
		_, _ = reader.ReadString('\n')
		pass, _ = reader.ReadString('\n')
		passwords := strings.Fields(pass)
		loginAttemp, _ = reader.ReadString('\n')
		loginAttemp = strings.TrimSuffix(loginAttemp, "\r\n")
		crack(passwords, loginAttemp)
		for k := range memo {
			delete(memo, k)
		}
		t--
	}

	//passwords := []string{"ab", "abcd", "cd"}
	//loginAttemp := "abcd"

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

			*res = slice[:len(slice)-1]

		}
	}
	return false

}
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
