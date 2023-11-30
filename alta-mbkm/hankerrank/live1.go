package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int64{87, 57, 370, 110, 90, 610, 2, 710, 140, 203, 150}

	fmt.Println(doubleSize(arr, 2))

	// fmt.Println(maxDifference([]int32{10, 5, 3, 1}))
	// fmt.Println(maxDifference2([]int32{5}))

	// fmt.Println(updateTimes([]int32{1, 2, 4, 4}, []int32{1, 2, 3, 4}))
	// fmt.Println(updateTimes2([]int32{1, 2, 4, 4}, []int32{1, 2, 3, 4}))

}

//* Problem 1
func doubleSize(arr []int64, b int64) int64 {
	//* copy the slice first then do the sorting
	// newArr := arr[:]
	// sort.Slice(newArr, func(i, j int) bool { return newArr[i] < newArr[j] })

	//* we can sort arr right away, no need to copy it into new slice
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	num := b
	for _, v := range arr {
		if v == num {
			num += v
		}
	}
	return num
}

//* Problem 2 - one test case did not pass
func maxDifference(px []int32) int32 {
	max := px[1] - px[0] // when len(px) == 1, will panic index out of range
	for i := 0; i < len(px); i++ {
		// this to handle out of index
		if i < len(px)-1 {
			if px[i+1] < px[i] {
				continue
			}
		}

		for j := i + 1; j < len(px); j++ {
			if px[j]-px[i] > max {
				max = px[j] - px[i]
			}
		}
	}
	if max < 0 {
		return -1
	}
	return max
}

//* Problem 2 - from kak robi (in question if all test case pass)
func maxDifference2(px []int32) int32 {
	var max int32 = -1
	min := px[0]
	for i, v := range px {
		if i == 0 {
			continue
		}

		if v < min {
			min = v
			continue
		}

		var count int32 = v - min
		if count > max {
			max = count
		}

	}
	return max
}


//* Problem 3
func updateTimes(signalOne []int32, signalTwo []int32) int32 {
	// Write your code here
	var res int32
	var maxequal int32
	if len(signalOne) < len(signalTwo) {
		for i := 0; i < len(signalOne); i++ {
			if signalOne[i] == signalTwo[i] {
				if signalOne[i] > maxequal {
					res += 1
					// maxequal = 0
					maxequal = signalOne[i]
				}
			}
		}
		return res
	}

	for i := 0; i < len(signalTwo); i++ {
		if signalTwo[i] == signalOne[i] {
			if signalTwo[i] > maxequal {
				res += 1
				// maxequal = 0
				maxequal = signalTwo[i]
			}
		}
	}
	return res
}

//* Problem 3 - Short version
func updateTimes2(signalOne []int32, signalTwo []int32) int32 {
	// Write your code here
	var res int32
	var maxequal int32
	var signals int
	if len(signalOne) < len(signalTwo) {
			signals = len(signalOne)
	} else {
		signals = len(signalTwo)
	}

	for i := 0; i < signals; i++ {
		if signalOne[i] == signalTwo[i] && signalOne[i] > maxequal {
			res += 1
			maxequal = signalOne[i]
		}
	}
	return res
}