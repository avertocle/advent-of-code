package main

import "fmt"

const SimTime = 175
const RC = 5
const CD = 2

var InitCulture = []int{1, 2, 3, 4, 1}

// func main() {
// 	culture := InitCulture
// 	for i := 0; i < SimTime; i++ {
// 		newCulture := make([]int, 0)
// 		for _, v := range currCulture {
// 			if v < RC {
// 				newCulture = append(newCulture, v+1)
// 			} else if v == RC {
// 				newCulture = append(newCulture, -1)
// 				newCulture = append(newCulture, -1)
// 			}
// 		}
// 		culture = simulate(culture)
// 		fmt.Printf("%v.", i)
// 	}
// 	fmt.Printf("%+v\n", culture)
// 	fmt.Printf("%v\n", len(culture))
// }

// func simulate(currCulture []int) []int {
// 	newCulture := make([]int, 0)
// 	for _, v := range currCulture {
// 		if v < RC {
// 			newCulture = append(newCulture, v+1)
// 		} else if v == RC {
// 			newCulture = append(newCulture, -1)
// 			newCulture = append(newCulture, -1)
// 		}
// 	}
// 	return newCulture
// }

func main() {
	culture := InitCulture
	for i := 0; i < SimTime; i++ {
		culture = simulate(culture)
		fmt.Printf("%v.", i)
	}
	//fmt.Printf("%+v\n", culture)
	fmt.Printf("\n\n%v\n", len(culture))
}

func simulate(currCulture []int) []int {
	newCulture := make([]int, 0)
	for _, v := range currCulture {
		if v < RC {
			newCulture = append(newCulture, v+1)
		} else if v == RC {
			newCulture = append(newCulture, -1)
			newCulture = append(newCulture, -1)
		}
	}
	return newCulture
}
