<<<<<<< HEAD
<<<<<<< HEAD
package main

import (
	"fmt"
	"math"
	"strings"
)

const prime = 101

func main() {
	//var t int
	//var pass, loginAttemp string
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Scanf("%d\n", &t)
	//for t > 0 {
	//	//	fmt.Scanln(&n)
	//	_, _ = reader.ReadString('\n')
	//	//tmp = strings.TrimSuffix(tmp, "\r\n")
	//	//n, _ := strconv.Atoi(tmp)
	//	//fmt.Scanf("%d\n", &n)
	//	pass, _ = reader.ReadString('\n')
	//	//pass = strings.TrimSuffix(pass, "\r\n")
	//	passwords := strings.Fields(pass)
	//	loginAttemp, _ = reader.ReadString('\n')
	//	loginAttemp = strings.TrimSuffix(loginAttemp, "\r\n")
	//	//	fmt.Printf("passwords: %#v \n", passwords)
	//	//	fmt.Printf("loginAttempt: %#v \n", loginAttemp)
	//	resolve(passwords, loginAttemp)
	//
	//	t--
	//}

	//naiveString("abcaaabcabaabcabacb", "ab")
	//rabinKarpMatcher("abcaaabcabaabcabacb", "abac", 256, 101)

	//KMP
	//kmpMatcher("abcabyabxabcabcaby", "abcaby")

	//dicer
	//diceRolls(3)

	diceSumRolls(3, 10)
}
func rabinKarpMatcher(text, pattern string, d, q int64) {
	tempPattern := []rune(pattern)
	tempText := []rune(text)
	n := len(text)
	m := len(pattern)
	h := int64(math.Pow(float64(d), float64(m-1)))
	var p, t int64
	end := m
	for i := 0; i < end; i++ {
		tempT := int64(tempText[i])
		tempP := int64(tempPattern[i])
		p = ((d * p) + tempP) % q
		t = ((d * t) + tempT) % q
	}
	for s := 0; s < n-m; s++ {
		if p == t {
			if pattern[:m] == text[s:s+m] {
				println("patter occurs with shift: ", s)
			}
		}
		if s < n-m {
			t = (d*(t-int64(tempText[s])*h) + int64(tempText[s+m])) % q
			if t < 0 {
				t = t + q
			}
		}
	}

}

func naiveString(text, pattern string) {
	n := len(text)
	m := len(pattern)
	end := n - m
	for s := 0; s < end; s++ {
		if pattern[:] == text[s:s+m] {
			println("patter occurs with shift: ", s)
		}

	}

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

/******************************Knuth-Morris-Pratt matching algorithm ******************************************/
func computePrefixFunction(pattern string) []int {
	p := []rune(pattern)
	m := len(p)
	π := make([]int, m)
	k := 0
	for q := 1; q < m; q++ {
		k = π[q-1]
		for k > 0 && p[k] != p[q] {
			k = π[k-1]
		}
		if p[k] == p[q] {
			k++
		}
		π[q] = k

	}

	return π

}

func kmpMatcher(text, pattern string) {
	t := []rune(text)
	p := []rune(pattern)
	n := len(t)
	m := len(p)
	π := computePrefixFunction(pattern)
	q := 0                   //number of characters matched
	for i := 0; i < n; i++ { //scan the text from left to right
		for q > 0 && p[q] != t[i] {
			q = π[q-1] //next character does not match
		}
		if p[q] == t[i] {
			q++ //next cbaracter matched
		}
		if q == m { //is all of P matched
			println("patter occurs with shift: ", (i+1)-m)
			q = π[q-1] //look the next match
		}

	}

}

/**********************PERMUTATION ALG*****************************/
func diceHelper(dice int, chosen []int) {
	if dice == 0 {
		//base case
		fmt.Println(chosen)

	} else {
		for i := 1; i <= 6; i++ {
			//chose
			chosen = append(chosen, i)
			//explore
			diceHelper(dice-1, chosen)
			// remove chosed
			//chosen = append(chosen[,],chosen[,]...)
			chosen = chosen[:len(chosen)-1]
		}
	}

}

func diceRolls(dice int) {
	var chosen []int
	diceHelper(dice, chosen)

}

/**********************PERMUTATION ALG*****************************/
func diceSumHelper(dice int, chosen []int, diseredSum, sumSofar int) {

	if dice == 0 {
		fmt.Println(chosen)

	} else {

		for i := 1; i <= 6; i++ {
			if sumSofar+i+1*(dice-1) <= diseredSum &&
				sumSofar+i+6*(dice-1) >= diseredSum {

				//chose
				chosen = append(chosen, i)
				//chosen = chosen[:len(chosen)+1]
				//chosen[len(chosen)-1] = i
				//explore
				diceSumHelper(dice-1, chosen, diseredSum, sumSofar+i)
				// remove chosed
				//chosen = append(chosen[,],chosen[,]...)
				//chosen = append(chosen[0:0], chosen[:len(chosen)-1]...)
				chosen = chosen[:len(chosen)-1]
			}
		}
	}

}

func diceSumRolls(dice int, desiredSum int) {
	var chosen []int
	//	chosen := make([]int, 0, dice)
	diceSumHelper(dice, chosen, desiredSum, 0)

}
=======
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
package main

import (
	"fmt"
	"math"
	"strings"
)

const prime = 101

func main() {
	//var t int
	//var pass, loginAttemp string
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Scanf("%d\n", &t)
	//for t > 0 {
	//	//	fmt.Scanln(&n)
	//	_, _ = reader.ReadString('\n')
	//	//tmp = strings.TrimSuffix(tmp, "\r\n")
	//	//n, _ := strconv.Atoi(tmp)
	//	//fmt.Scanf("%d\n", &n)
	//	pass, _ = reader.ReadString('\n')
	//	//pass = strings.TrimSuffix(pass, "\r\n")
	//	passwords := strings.Fields(pass)
	//	loginAttemp, _ = reader.ReadString('\n')
	//	loginAttemp = strings.TrimSuffix(loginAttemp, "\r\n")
	//	//	fmt.Printf("passwords: %#v \n", passwords)
	//	//	fmt.Printf("loginAttempt: %#v \n", loginAttemp)
	//	resolve(passwords, loginAttemp)
	//
	//	t--
	//}

	//naiveString("abcaaabcabaabcabacb", "ab")
	//rabinKarpMatcher("abcaaabcabaabcabacb", "abac", 256, 101)

	//KMP
	//kmpMatcher("abcabyabxabcabcaby", "abcaby")

	//dicer
	//diceRolls(3)

	diceSumRolls(3, 10)
}
func rabinKarpMatcher(text, pattern string, d, q int64) {
	tempPattern := []rune(pattern)
	tempText := []rune(text)
	n := len(text)
	m := len(pattern)
	h := int64(math.Pow(float64(d), float64(m-1)))
	var p, t int64
	end := m
	for i := 0; i < end; i++ {
		tempT := int64(tempText[i])
		tempP := int64(tempPattern[i])
		p = ((d * p) + tempP) % q
		t = ((d * t) + tempT) % q
	}
	for s := 0; s < n-m; s++ {
		if p == t {
			if pattern[:m] == text[s:s+m] {
				println("patter occurs with shift: ", s)
			}
		}
		if s < n-m {
			t = (d*(t-int64(tempText[s])*h) + int64(tempText[s+m])) % q
			if t < 0 {
				t = t + q
			}
		}
	}

}

func naiveString(text, pattern string) {
	n := len(text)
	m := len(pattern)
	end := n - m
	for s := 0; s < end; s++ {
		if pattern[:] == text[s:s+m] {
			println("patter occurs with shift: ", s)
		}

	}

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

/******************************Knuth-Morris-Pratt matching algorithm ******************************************/
func computePrefixFunction(pattern string) []int {
	p := []rune(pattern)
	m := len(p)
	π := make([]int, m)
	k := 0
	for q := 1; q < m; q++ {
		k = π[q-1]
		for k > 0 && p[k] != p[q] {
			k = π[k-1]
		}
		if p[k] == p[q] {
			k++
		}
		π[q] = k

	}

	return π

}

func kmpMatcher(text, pattern string) {
	t := []rune(text)
	p := []rune(pattern)
	n := len(t)
	m := len(p)
	π := computePrefixFunction(pattern)
	q := 0                   //number of characters matched
	for i := 0; i < n; i++ { //scan the text from left to right
		for q > 0 && p[q] != t[i] {
			q = π[q-1] //next character does not match
		}
		if p[q] == t[i] {
			q++ //next cbaracter matched
		}
		if q == m { //is all of P matched
			println("patter occurs with shift: ", (i+1)-m)
			q = π[q-1] //look the next match
		}

	}

}

/**********************PERMUTATION ALG*****************************/
func diceHelper(dice int, chosen []int) {
	if dice == 0 {
		//base case
		fmt.Println(chosen)

	} else {
		for i := 1; i <= 6; i++ {
			//chose
			chosen = append(chosen, i)
			//explore
			diceHelper(dice-1, chosen)
			// remove chosed
			//chosen = append(chosen[,],chosen[,]...)
			chosen = chosen[:len(chosen)-1]
		}
	}

}

func diceRolls(dice int) {
	var chosen []int
	diceHelper(dice, chosen)

}

/**********************PERMUTATION ALG*****************************/
func diceSumHelper(dice int, chosen []int, diseredSum, sumSofar int) {

	if dice == 0 {
		fmt.Println(chosen)

	} else {

		for i := 1; i <= 6; i++ {
			if sumSofar+i+1*(dice-1) <= diseredSum &&
				sumSofar+i+6*(dice-1) >= diseredSum {

				//chose
				chosen = append(chosen, i)
				//chosen = chosen[:len(chosen)+1]
				//chosen[len(chosen)-1] = i
				//explore
				diceSumHelper(dice-1, chosen, diseredSum, sumSofar+i)
				// remove chosed
				//chosen = append(chosen[,],chosen[,]...)
				//chosen = append(chosen[0:0], chosen[:len(chosen)-1]...)
				chosen = chosen[:len(chosen)-1]
			}
		}
	}

}

func diceSumRolls(dice int, desiredSum int) {
	var chosen []int
	//	chosen := make([]int, 0, dice)
	diceSumHelper(dice, chosen, desiredSum, 0)

}
<<<<<<< HEAD
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
