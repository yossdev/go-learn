package main

import (
	"fmt"
	"sort"
)

func main() {
	//* Problem 1 - Simple Equations
	SimpleEquations(1, 2, 3)	// no solution
	SimpleEquations(6, 6, 14)	// 1, 2, 3
	SimpleEquations(679, 5608800, 233065)	// 100, 123, 456 (custom case)

	//* Problem 2 - Money Coins
	fmt.Println(moneyCoins(123))   // [100 20 1 1 1]
	fmt.Println(moneyCoins(432))   // [200 200 20 10 1 1]
	fmt.Println(moneyCoins(543))   // [500, 20, 20, 1, 1, 1]
	fmt.Println(moneyCoins(7752))  // [5000, 2000, 500, 200, 50, 1, 1]
	fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]

	//* Problem 3 - Dragon of Loowater
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    // 11
	DragonOfLoowater([]int{5, 10}, []int{5})         // knight fall
 	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall
 	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10
	
	//* Problem 4 - Binary Search Algorithm
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3) // 2
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5) // 3
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 69, 60}, 53)  // 6
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1
}

//* Problem 1 - Simple Equations
func SimpleEquations(a, b, c int) {
	for x := 1; x < a; x++ {
		for y := x + 1; y < a+1; y++ {
			for z := y; z < a+1; z++ {
				if x*y*z == b && x+y+z == a {
					fmt.Println(x,y,z)
					return
				}
			}
		}
	}
	fmt.Println("no solution")
}


//* Problem 2 - Money Coins
func moneyCoins(money int) []int {
	coins := [11]int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	var n []int 
	var i int
	if money <= coins[4] {
		i = 4
	} else {
		i = 0
	}

	for j := i; j <= len(coins); j++ {
		if money >= coins[i] {
			i = j
			break
		}
	}

	for {
		if money >= coins[i] {
			money -= coins[i]
			n = append(n, coins[i])
		}
		if money - coins[i] < 0 {
			i++
		}
		if money == 0 {
			break
		}
	}
	return n
}


//* Problem 3 - Dragon of Loowater
func DragonOfLoowater(dragonHead, knightHeight []int) {
	if len(dragonHead) > len(knightHeight) {
		fmt.Println("knight fall")
		return
	}

	sort.Ints(dragonHead)
	sort.Ints(knightHeight)

	if dragonHead[len(dragonHead) - 1] > knightHeight[len(knightHeight) - 1] {
		fmt.Println("knight fall")
		return
	}

	var i int
	var j int
	var sum int
	for {
		if dragonHead[i] <= knightHeight[j] {
			sum += knightHeight[j]
			i++
		}
		if i == len(dragonHead) {
			break
		}
		j++
	}
	fmt.Println(sum)
}


//* Problem 4 - Binary Search Algorithm
func BinarySearch(array []int, x int) {
	var o, l, r int = 0, 1, len(array) - 1
	for l <= r && o == 0 {
		m := (l+r) / 2
		if x < array[m] {
			r = m - 1
		} else if x > array[m] {
			l = m + 1
		} else {
			o = m
		}
	}
	if o == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(o)
	}
}