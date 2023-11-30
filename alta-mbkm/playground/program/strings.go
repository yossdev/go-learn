package program

import (
	"fmt"
	"strings"
)

func PlayWithString() {
	// //* 1. Contain func
	str := "haloalala "
	subStr := " "
	contain := strings.Contains(str, subStr)
	fmt.Println(contain)

	// //* 2. Replace
	// //* -1 to replace all, 1 - n will replace n first
	// replace := strings.Replace(str, "a", "i", -1)
	// replace2 := strings.Replace(str, "a", "i", 3)
	// fmt.Println(replace)
	// fmt.Println(replace2)

	// //* 3. Compare
	// str2 := "halo"
	// fmt.Println(str == str2)

	// //* 4. Get string value at n index
	// getStr := str[1:len(str)]
	// fmt.Println(getStr)

	// //* 5. Insert to a string
	// insert := "HaloSemua"
	// index := 4
	// res := insert[:index] + " Hai " + insert[index:]
	// fmt.Println(insert, res)

	// //TODO 6. Buat Function String to array
}