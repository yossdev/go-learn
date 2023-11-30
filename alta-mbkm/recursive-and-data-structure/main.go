package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	//* Problem 1 done
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29

	//* Problem 2 done
	fmt.Println(fibonacci(0))  // 0
   	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144

	//* Problem 3 done
	primaSegiEmpat(2, 3, 13)
	/*
    17 19
    23 29
    31 37
    156
	*/
	primaSegiEmpat(5, 2, 1)
	/*
    2  3  5  7 11
    13 17 19 23 29
    129
	*/

	//* Problem 4 done
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))	// 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))	// 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))	// 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))	// 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))		// 12
	fmt.Println(MaxSequence([]int{2, 5, 6, 2, 3, 1, 6, 6}))
	fmt.Println(MaxSequence([]int{-2, -5, -6, -2, -3, -1, -6, -6}))

	//* Problem 5 done
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	// min: -2 index: 3 max: 8 index: 5

	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	// min: -5 index: 1 max: 22 index: 3

	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	// min: -21 index: 4 max: 9 index: 2

	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	// min: -1 index: 0 max: 18 index: 5

	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
	// min: -20 index: 5 max: 7 index: 4

	//* Problem 6 done
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})      // 3 
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4 
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // 4
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // 1
	MaximumBuyProduct(0, []int{10000, 30000})                        // 0

	//* Problem 7 done
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	// [3, 4]

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	// [6 5]

	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
	// “tutup kartu”

	//* Problem 8 done
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang->1 ruby->2 js->4

	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C->1 D->2 B->3 A->4

	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// football->1 basketball->1 tenis->1
}

//* Problem 1 - Prima ke X
func primeX(number int) int {
	if number < 1 {
		return 0
	}
	i := 2
	for {
		isPrime := checkPrime(i)
		if isPrime {
			number--
		}
		if number == 0 {
			return i
		}
		i++
	}
}

func checkPrime(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number % i == 0  {
			return false
		}
	}
	return true
}


//* Problem 2 – Fibonacci (Recursive)
func fibonacci(number int) int {
	if number <= 0 {
		return 0
	}
	if number < 2 {
		return number
	}
	return fibonacci(number - 1) + fibonacci(number - 2)
}


//* Problem 3 – Prima Segi Empat
func primaSegiEmpat(high, wide, start int) {
	primes := genPrime(start, (high * wide))

	sum := 0
	for _, value := range primes {
		sum += value
	}

	count := -1
	for i := 0; i < wide; i++ {
		for j := 0; j < high; j++ {
			if start < 6 && count < 5 {
				fmt.Print(" ")
			}
			count += 1
			fmt.Print(primes[count], " ")
		}
		fmt.Println()
	}
	fmt.Println(sum)
}

func genPrime(start, number int) []int {
	var res []int
	i := start + 1
	for {
		isPrime := checkPrime(i)
		if isPrime {
			res = append(res, i)
			number--
		}
		if number == 0 {
			return res
		}
		i++
	}
}

//* Problem 4 - Total Maksimum Dari Deret Bilangan
func MaxSequence(arr []int) int {
	plus, isPositive := allPositive(arr)
	if isPositive {
		return plus
	}

	sum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			for j := i; j < len(arr); j++ {
				sum += arr[j]
				if sum < 0 {
					sum = 0
					break
				} 
				if arr[i] + arr[j] <= 1 {
					sum += (arr[j] * -1)
					return sum
				}
			}
		}
	}
	
	if sum == 0 {
		sort.Ints(arr)
		return arr[len(arr) - 1]
	}

	return sum
}

func allPositive(bil []int) (int, bool) {
	positive := 0
	if bil[0] > 0 {
		for i := 0; i < len(bil); i++ {
			if bil[i] < 0 {
				return 0, false
			}
			positive += bil[i]
		}
		return positive, true
	}
	return 0, false
}


//* Problem 5 - Find Min and Max Number
func FindMinAndMax(arr []int) string {
	resMin := 0
	min := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < resMin {
			min = 0
			min += i
			resMin = arr[i]
			
		}
	}

	resMax := 0
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > resMax {
			max = 0
			max += i
			resMax = arr[i]
			
		}
	}
	return "min: " + strconv.Itoa(resMin) + " index: " + strconv.Itoa(min) + " max: " + strconv.Itoa(resMax) + " index: " + strconv.Itoa(max)
}


//* Problem 6 - Maximum Buy Product
func MaximumBuyProduct(money int, productPrice []int) {
	bought := 0
	sort.Ints(productPrice)
	for _, v := range productPrice {
		if money - v < 0 {
			break
		}
		money -= v
		bought += 1
	}
	fmt.Println(bought)
}


//* Problem 7 - Playing Domino
func playingDomino(cards [][]int, deck []int) interface{} {
	for i, _ := range cards {
		if cards[i][0] == deck[0] || cards[i][0] == deck[1] {
			if cards[i][0] + cards[i][1] >= deck[0] + deck[1] {
				return cards[i]
			}
		}
	}

	for j, _ := range cards {
		if cards[j][1] == deck[0] || cards[j][1] == deck[1] {
			if cards[j][0] + cards[j][1] >= deck[0] + deck[1] {
				return cards[j]
			}
		}
	}
	return "tutup kartu"
}




//* Problem 8 - Most Appear Item
type pair struct {
	items string
	n int
}

func MostAppearItem(items []string) []pair {
	collection := make(map[string]int)
	for i := 0; i < len(items); i++ {
		itemsFound := searchItems(items, items[i])
		if itemsFound {
			collection[items[i]] += 1
		}
	}

	res := make([]pair, 0)
	for j := 1; j < nMax(collection) + 1; j++ {
		for i, v := range collection {
			if v == j {
				res = append(res, pair{i, v})
			}
		}
	}
    return res
}

func searchItems(items []string, value string) bool {
	for _, s := range items {
		if s == value {
			return true
		}
	}
	return false
}

func nMax(n map[string]int) int {
	max := 0
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	return max
}