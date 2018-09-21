package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var totScoreA, totScoreB int
	reader := bufio.NewReader(os.Stdin)
	in, _ := reader.ReadString('\n')
	scoreA := strings.Fields(in)
	in, _ = reader.ReadString('\n')
	scoreB := strings.Fields(in)

	for i := 0; i < len(scoreA); i++ {

		tmpA, _ := strconv.Atoi(scoreA[i])
		tmpB, _ := strconv.Atoi(scoreB[i])
		if tmpA > tmpB {
			totScoreA++
		}
		if tmpA < tmpB {
			totScoreB++
		}
	}

	fmt.Println(totScoreA, totScoreB)

}
