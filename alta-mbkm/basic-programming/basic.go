package main

import (
	"fmt"
)

func main() {
	// luasTabung()

	// gradeNilai()

	// faktorBil()

	// primeNumber(2)
	// primeNumber(13)
	// primeNumber(17)
	// primeNumber(20)
	// primeNumber(35)
	// primeNumber(3000000007)

	// palindrome("civic")
	// palindrome("katak")
	// palindrome("kasur rusak")
	// palindrome("mistar")
	// palindrome("lion")

	fmt.Println(pangkat(10000000, 5))
	// fmt.Println(pangkat(5, 3))
	// fmt.Println(pangkat(10, 2))
	// fmt.Println(pangkat(2, 5))
	// fmt.Println(pangkat(7, 3))

	// playWithAsterik(5)

	// cetakTabelPerkalian(30)
	// cetakTabelPerkalian(8)
}

//* Problem 1 - Menghitung Luas Permukaan Tabung
func luasTabung() {
	//* input hard coded
	// T := 20.0
	// r := 4.0

	//* Next Problem: input dari user
	fmt.Printf("Enter T value: ")
	var T float64
	fmt.Scanf("%g\n", &T)
	// fmt.Println(T)

	fmt.Printf("Enter r value: ")
	var r float64
	fmt.Scanf("%g\n", &r)
	// fmt.Println(r)

	// kode disini
	const Pi = 3.14
	var res float64 = 2.0 * Pi * r * (r + T)
	fmt.Println(res)
}

//* Problem 2 - Grade Nilai
func gradeNilai() {
	var studentScore int = 80
	switch {
	case studentScore >= 80 && studentScore <= 100:
		fmt.Println("A")
	case studentScore >= 65 && studentScore <= 79:
		fmt.Println("B")
	case studentScore >= 50 && studentScore <= 64:
		fmt.Println("C")
	case studentScore >= 35 && studentScore <= 49:
		fmt.Println("D")
	case studentScore >= 0 && studentScore <= 34:
		fmt.Println("E")
	default:
		fmt.Println("Nilai Invalid")
	}
}

//* Problem 3 - Faktor Bilangan
func faktorBil() {
	var bilangan int
	fmt.Scanf("%d", &bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan % i == 0 {
			fmt.Println(i)
		}
	}
}

//* Problem 4 - Bilangan Prima
func primeNumber(number int) bool {
    switch {
    case number == 1:
		fmt.Println("Bukan Bilangan Prima")
        return false
    case number > 1:
        for i := 2; i < number; i++ {
            if number % i == 0  {
				fmt.Println("Bukan Bilangan Prima")
				fmt.Println(i)
                return false
            }
        }
	default:
		fmt.Println("Bukan Bilangan Prima")
		return false
    }
	fmt.Println("Bilangan Prima")
	return true
}

//* Problem 5 - Palindrome
func palindrome(input string) bool {
	runeStr := input
	var res string

	for i := 0; i < len(runeStr); i++ {
		res += string(runeStr[len(runeStr) - (1 + i)])
	}
	
	if input == res {
		fmt.Println("Palindrome")
		return true
	}
	fmt.Println("Bukan palindrome")
	return false
}

//* Problem 6 - Exponentiation
func pangkat(base, pangkat int) int {
	var res int = 1
	for i := 0; i < pangkat; i++ {
		fmt.Println(res)
		res *= base
	}
	fmt.Println("1")
	return res
}

//* Problem 7 - Play With Asterisk
func playWithAsterik(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n - i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	} 
}

//* Problem 8 - Cetak Tabel Perkalian
func cetakTabelPerkalian(number int) {
	var base int
	// fmt.Printf("Banyak Tabel: ")
	// fmt.Scanf("%d\n", &number)
	spacing := "  "
	maxDigitLength := len(fmt.Sprintf("%d", (number * number)))

	for i := 1; i <= number; i++ {
		for j := 0; j < number; j++ {
			base += i

			currDigitLength := len(fmt.Sprintf("%d", base))
			fmt.Printf("%s%s%d", spacing, strings.Repeat(" ", maxDigitLength-currDigitLength), base)
		}
		base = 0
		fmt.Println()
	}
}
