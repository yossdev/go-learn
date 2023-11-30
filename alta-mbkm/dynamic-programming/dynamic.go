package main

func main() {
	//* Problem 1 - Fibonacci Number Top-down
	// fmt.Println(fibo(0))  // 0
	// fmt.Println(fibo(1))  // 1
	// fmt.Println(fibo(2))  // 1
	// fmt.Println(fibo(3))  // 2
	// fmt.Println(fibo(5))  // 5
	// fmt.Println(fibo(6))  // 8
	// fmt.Println(fibo(7))  // 13
	// fmt.Println(fibo(9))  // 34
	// fmt.Println(fibo(10)) // 55

	//* Problem 2 - Fibonacci Number Bottom-up
	// fmt.Println(fibo(0))  // 0
	// fmt.Println(fibo(1))  // 1
	// fmt.Println(fibo(2))  // 1
	// fmt.Println(fibo(3))  // 2
	// fmt.Println(fibo(5))  // 5
	// fmt.Println(fibo(6))  // 8
	// fmt.Println(fibo(7))  // 13
	// fmt.Println(fibo(9))  // 34
	// fmt.Println(fibo(10)) // 55

	//* Problem 3 - Frog
	// fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	// fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}

//* Problem 1 - Fibonacci Number Top-down (Recursive with memoization) from Frendi Gunawan
func fibo(n int) int {
	// TODO: To be finish
	return 1
}

//* Problem 1 - Fibonacci Number Top-down (Recursive with memoization)
// func fibo(n int) int {
// 	switch {
// 	case n <= 0:
// 		return 0
// 	case n <= 2:
// 		return 1
// 	}

// 	cache := make(map[int]int)
//     bucket := make([]int, n)

//     for i := 1; i <= n; i++{
// 		bucket[i-1] = memoFib(i, cache)
// 	}
// 	result := bucket[n-1]
// 	return result
// }

// func memoFib(n int, cache map[int]int) int {
// 	if n < 2 {
// 		cache[n] = n
// 		return n
// 	}
// 	if _, ok := cache[n-1]; !ok {
// 		cache[n-1] = memoFib(n-1, cache)
// 	}
// 	if _, ok := cache[n-2]; !ok {
// 		cache[n-2] = memoFib(n-2, cache)
// 	}
// 	return cache[n-1] + cache[n-2]
// }

// * Problem 2 - Fibonacci Number Bottom-up (For loop with out Tabulation)
// func fibo(n int) int {
// 	prev := 0
// 	res := 1
// 	switch {
// 	case n <= 0:
// 		return prev
// 	case n <= 2:
// 		return res
// 	}

// 	for i := 1; i < n; i++ {
// 		sum := prev + res
// 		prev = res
// 		res = sum
// 	}
// 	return res
// }

// * Problem 2 - Fibonacci Number Bottom-up (For loop with Tabulation)
// func fibo(n int) int {
// 	var res = []int{0, 1}
// 	switch {
// 	case n <= 0:
// 		return res[0]
// 	case n <= 2:
// 		return res[1]
// 	}

// 	for i := 2; i <= n; i++ {
// 		res = append(res, (res[i-1] + res[i-2]))
// 	}
// 	return res[n]
// }

//* Problem 3 - Frog
// func Frog(jumps []int) int {
// 	if len(jumps) < 4 {
// 		if len(jumps) == 2 {
// 			return int(math.Abs(float64(jumps[0] - jumps[1])))
// 		}
// 		if math.Abs(float64(jumps[0] - jumps[1])) < math.Abs(float64(jumps[0] - jumps[2])) {
// 			return int(math.Abs(float64(jumps[0] - jumps[1])))
// 		}
// 		return int(math.Abs(float64(jumps[0] - jumps[2])))
// 	}

// 	cost1 := math.Abs(float64(jumps[1] - jumps[0]))
// 	cost2 := math.Abs(float64(jumps[2] - jumps[0]))
// 	var c []int

// 	// cost1
// 	i := 1
// 	for {
// 		for j := i+1; j <= i+2; j++ {
// 			c = append(c, int(math.Abs(float64(jumps[i] - jumps[j]))))
// 			if j == len(jumps) - 1 {
// 				break
// 			}
// 		}
// 		if len(c) == 2 {
// 			if c[0] < c[1] {
// 				cost1 += float64(c[0])
// 				i += 1
// 			} else {
// 				cost1 += float64(c[1])
// 				i += 2
// 			}
// 		} else {
// 			cost1 += float64(c[0])
// 			i += 1
// 		}
// 		if i == len(jumps) - 1 {
// 			break
// 		}
// 		c = nil
// 	}

// 	// cost2
// 	c = nil
// 	k := 2
// 	for {
// 		for l := k+1; l <= k+2; l++ {
// 			c = append(c, int(math.Abs(float64(jumps[k] - jumps[l]))))
// 			if l == len(jumps) - 1 {
// 				break
// 			}
// 		}
// 		if len(c) == 2 {
// 			if c[0] < c[1] {
// 				cost2 += float64(c[0])
// 				k += 1
// 			} else {
// 				cost2 += float64(c[1])
// 				k += 2
// 			}
// 		} else {
// 			cost2 += float64(c[0])
// 			k += 1
// 		}
// 		if k == len(jumps) - 1 {
// 			break
// 		}
// 		c = nil
// 	}

// 	if cost1 < cost2 {
// 		return int(cost1)
// 	}
// 	return int(cost2)
// }