package program

import "fmt"

//* Tugas Algo Bil. Prima
func PrimeNum() string {
    var input int
    fmt.Scanln(&input)
    switch {
    case input == 2 || input == 3:
        return "Bilangan Prima"
    case input > 3:
        for i := 1; i <= input; i++ {
            /* prime number can be written in the form of 6n + 1 or 6n – 1
            (except the multiples of prime numbers, i.e. 2, 3, 5, 7, 11)*/
            primeOne := 6 * i - 1
            primeTwo := 6 * i + 1
            if input == primeOne || input == primeTwo {
                return "Bilangan Prima"
            }
        }
    }
    return "Bukan Bilangan Prima"
}

