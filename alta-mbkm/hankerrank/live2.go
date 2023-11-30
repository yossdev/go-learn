package main

import (
	"sort"
)

// func main() {
// 	fmt.Println(funWithAnagrams([]string{"code", "aaagmnrs", "anagrams", "doce"}))
// }

//* Price Check
func priceCheck(products []string, productPrices []float32, productSold []string, soldPrice []float32) int32 {
    // Write your code here
    var error int32
    for i := 0; i < len(productSold); i++ {
        for j := 0; j < len(products); j++ {
            if productSold[i] == products[j] {
                if soldPrice[i] == productPrices[j] {
                    break
                } else {
                    error += 1
                }
            }
        }
    }
    return error
}

//* Maximum Rating Sum (of sub-array) - The tricky part is in data type like int, int32, int64, etc.
func maximumSum(arr []int32) int64 {
    // Write your code here
    tab := make([]int64, len(arr))
	var res int64 = int64(arr[0])
    tab[0] = int64(arr[0])
    for i := 1; i < len(arr); i++ {
        if tab[i-1] > 0 {
            tab[i] = int64(arr[i]) + int64(tab[i-1])
        } else {
            tab[i] = int64(arr[i])
        }
        if res < tab[i] {
            res = tab[i]
        }
    }
    return res
}

//* Fun with Anagrams
func funWithAnagrams(text []string) []string {
    // Write your code here
    var res = []string{}
    var str = [][]string{}
    for i := 0; i < len(text); i++ {
        dict := dict(text[i])
        sort.Strings(dict)
        str = append(str, dict)
    }

    for j := 0; j < len(str); j++ {
        for k := j+1; k < len(str); k++ {
            if toStr(str[j]) == toStr(str[k]) {
                text[k] = "z"
                break
            }
        }
    }
    sort.Strings(text)
    for _, v := range text {
        if v != "z" {
            res = append(res, v)
        }
    }
    return res
}

func dict(s string) []string {
    dict := make([]string, 0)
    for i := 0; i < len(s); i++ {
        dict = append(dict, string(s[i]))
    }
    return dict
}

func toStr(str []string) string {
    var newStr string
    for _, v := range str {
        newStr += v
    }
    return newStr
}