package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var hour string
	//fmt.Scanf("%v\n", &hour)
	hour = "12:40:22AM"
	//fmt.Println(hour)
	hour = strings.ToUpper(hour)
	ind := strings.Index(hour, "PM")
	if ind != -1 {

		tmp := hour[:ind]
		//	fmt.Println(tmp)
		prefS := hour[:2]
		num, _ := strconv.Atoi(prefS)
		//tmp1 := tmp[2:]
		if num < 12 {
			num = num + 12
		}
		fmt.Printf("%v%s\n", num, tmp)
	} else {
		ind := strings.Index(hour, "AM")
		tmp := hour[:ind]
		prefS := hour[:2]
		num, _ := strconv.Atoi(prefS)
		//tmp1 := tmp[2:]
		if num == 12 {

			fmt.Printf("00%s\n", tmp[2:])
		} else {

			fmt.Printf("%v\n", tmp)
		}

	}
}
