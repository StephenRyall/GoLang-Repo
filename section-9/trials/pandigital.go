// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func Ints(input []int) []int {
// 	u := make([]int, 0, len(input))
// 	m := make(map[int]bool)

// 	for _, val := range input {
// 		if _, ok := m[val]; !ok {
// 			m[val] = true
// 			u = append(u, val)
// 		}
// 	}

// 	return u
// }

// func main() {
// 	u := make([]string, 0)

// 	for s := 1; s < 99999; s++ {
// 		for b := 1; b < 99999; b++ {
// 			p := s * b
// 			pToString := strconv.FormatInt(int64(p), 10)
// 			pToArray := strings.Split(pToString, "")
// 			pFinal := strconv.ParseInt(pToArray, 10, 0)
// 			sToString := strconv.FormatInt(int64(s), 10)
// 			sToArray := strings.Split(sToString, "")
// 			bToString := strconv.FormatInt(int64(b), 10)
// 			bToArray := strings.Split(bToString, "")
// 			stringArray1 := append(pToArray, sToArray...)
// 			totalStringArray := append(stringArray1, bToArray...)

// 			if len(totalStringArray) >= 10 {
// 				break
// 			}
// 			if len(totalStringArray) == 9 {
// 				for _, v := range totalStringArray {
// 					if v == "0" {
// 						continue
// 					}
// 				}
// 				totalArray, _ := strconv.ParseInt(totalStringArray, 10, 0)

// 			}
// 		}
// 	}

// 	fmt.Println(u)
// }
