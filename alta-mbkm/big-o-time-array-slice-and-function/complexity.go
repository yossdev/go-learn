package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// primeNumber(1000000007)
	// primeNumber(1500450271)
	// primeNumber(1000000000)
	// primeNumber(10000000019)
	// primeNumber(10000000033)

	// fmt.Println(pow(10, 5))
	// fmt.Println(pow(2, 3))
	// fmt.Println(pow(7, 2))
	// fmt.Println(pow(17, 6))
	// fmt.Println(pow(5, 3))

	
	// fmt.Println(pow2(2, 5))
	// fmt.Println(pow2(2, 4))

	defer elapsed("page")()  // <-- The trailing () is the deferred call
    // fmt.Println(pow2(17, 6))
	fmt.Println(pow(10, 500))
	time.Sleep(time.Second * 2)

	// fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// fmt.Println(ArrayMerge([]string{"devil jin"}, []string{"sergei"}))
	// fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// fmt.Println(ArrayMerge([]string{}, []string{}))
	
	// fmt.Println(munculSekali("1234123"))
	// fmt.Println(munculSekali("76523752"))
	// fmt.Println(munculSekali("12345"))
	// fmt.Println(munculSekali("1122334455"))
	// fmt.Println(munculSekali("0872504"))

	// fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	// fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	// fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	// fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	// fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}

func elapsed(what string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", what, time.Since(start))
    }
}

//* Problem 1 - Bilangan Prima
func primeNumber(number int) bool {
	switch {
	case number == 1 || number < 1:
		fmt.Println("Bukan Bilangan Prima")
		return false
	case number == 2 || number == 3 || number == 5 || number == 7:
		fmt.Println("Bilangan Prima")
		return true
	case number % 2 == 0 || number % 3 == 0 || number % 5 == 0 || number % 7 == 0:
		fmt.Println("Bukan Bilangan Prima")
		return false
	default:
		for i := 2; i < number; i++ {
			if (number % i == 0)  {
				fmt.Println("Bukan Bilangan Prima")
				return false
			}
		}
	}
	fmt.Println("Bilangan Prima")
	return true
}


//* Problem 2 - Fast Exponentiation (Recursive)
func pow(x, n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 1/pow(x,-n)
	} else if n % 2 == 0 {
		return pow(x, n/2) * pow(x, n/2)
	}
	return x*pow(x, (n-1)/2)*pow(x, (n-1)/2)
}

//* Problem 2 - Fast Exponentiation (Divide and Conquer)
func pow2(x, n int) int {
	base := x
	pow := n 
	if n == 0 {
		return 1
	} else {
		n = n / 2 - 1
	}

	for i := 0; i < n; i++ {
		// fmt.Println(i, x)
		x *= base
	}

	if pow % 2 == 1 {
		return x * x * base
	}

	return x * x
}


//* Problem 3 - Array Merge
func ArrayMerge(arrayA, arrayB []string) []string {
	res := make([]string, len(arrayA))
	copy(res, arrayA)
	check := 0
	for i := 0; i < len(arrayB); i++{
		for j := 0; j < len(arrayA); j++ {
			if arrayB[i] == arrayA[j] {
				check += 1
			} else {
				check += 0
			}
		}
		if check == 0 {
			res = append(res, arrayB[i])
		}
		check = 0
	}
	return res
}


//* Problem 4 - Angka Muncul Sekali
func munculSekali(angka string) []int  {
	var res int
	arr := make([]int, 0)
	for i := 0; i < len(angka); i++ {
		for j := 0; j < len(angka); j++ {
			if angka[i] == angka[j] {
				res += 1
			}
		}
		if res == 1 {
			arrInt, _ := strconv.Atoi(string(angka[i]))
			arr = append(arr, arrInt)
		}
		res = 0
	}
	return arr
}


//* Problem 5 - Pair with Target Sum
func PairSum(arr []int, target int) []int {
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] + arr[j] == target {
				res = append(res, i)
				res = append(res, j)
				return res
			}
		}
	}
	return res
}