package main

import (
	"fmt"
)

func aryUnique(arr []int) map[int]int {
	m := make(map[int]int)
	for _, s := range arr {

		m[s] += 1
	}
	return m
}

func RemoveIndex(a []int, index int) []int {
	return append(a[:index], a[index+1:]...)
}

func findLeastNumOfUniqueInts(arr []int, k int) int {

	if k == 0 {
		m := aryUnique(arr)
		return len(m)
	}

	m := aryUnique(arr)
	t := 0
	n := 1
	l := 0
	fmt.Println(m)

	if len(m) == 1 {
		for l < k {
			arr = RemoveIndex(arr, 0)
			l++
		}

		m = aryUnique(arr)
	} else {
	Uniq:
		for t < k {

			for key, val := range m {

				// 削除するelement数に達してたら
				//
				if t == k || k-t < n {
					break Uniq
				}

				if val == n {
					delete(m, key)
					t += val

				}
				// fmt.Println("t=")
				// fmt.Println(t)
				// fmt.Println("n=")
				// fmt.Println(n)
				// fmt.Println(n)

				// indexを見つけて、arrからも削除
				// index := pos(arr, key)
				// arr = RemoveIndex(arr, index)
				// fmt.Println("k-t=")
				// fmt.Println(k - t)

				fmt.Printf("k-t: %d, n:%d\n", k-t, n) // Prints `Binary: 100\101`

			}

			n++
		}

	}

	return len(m)
}

func main() {
	fmt.Println("-----------start leetcode1481--------------")

	// arr := []int{2, 1, 1, 3, 3}
	// arr := []int{24, 119, 157, 446, 251, 117, 22, 168, 374, 373, 323, 311, 441, 213, 120, 412, 200, 236, 328, 24, 164, 104, 331, 32, 19, 223, 89, 114, 152, 82, 456, 381, 355, 343, 157, 245, 443, 368, 229, 49, 82, 16, 373, 142, 240, 125, 8}
	arr := []int{4, 3, 1, 1, 3, 3, 2}

	a := findLeastNumOfUniqueInts(arr, 3)
	// a := findLeastNumOfUniqueInts(arr, 41)

	fmt.Println(a)

	fmt.Println("-----------end leetcode1481--------------")
}
