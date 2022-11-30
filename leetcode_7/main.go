package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {

	str := strconv.Itoa(x)

	// fmt.Println(str)

	var strAry []string

	for _, c := range str {
		strAry = append(strAry, string(c))
	}

	isMinus := false
	if strAry[0] == "-" {
		isMinus = true
		strAry = strAry[1:]

	}

	for i, j := 0, len(strAry)-1; i < j; i, j = i+1, j-1 {
		strAry[i], strAry[j] = strAry[j], strAry[i]
	}

	if strAry[0] == "0" {
		strAry = strAry[1:]
	}

	reverseStr := strings.Join(strAry[:], "")

	// fmt.Println(reverseStr)

	reverseInt, err := strconv.Atoi(reverseStr)
	if err != nil {
		fmt.Println("Error during conversion")
		return 0
	}

	// もし-なら
	if isMinus {
		reverseInt = 0 - reverseInt

		x := math.Pow(-2, 31)

		if int(x) > reverseInt {
			reverseInt = 0
		}
	} else {

		x := math.Pow(2, 31)

		if int(x) < reverseInt {
			reverseInt = 0
		}
	}

	return reverseInt

}

func main() {

	fmt.Println(reverse(213))
	fmt.Println(reverse(-213))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-1534236469))
}
