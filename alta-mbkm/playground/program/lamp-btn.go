package program

import "fmt"

func LampBtn() {
	var N int
	fmt.Scanln(&N)
	btnPressed := 0
	for i := 1; i <= N; i++ {
		if N % i == 0 {
			btnPressed += 1
		}
	}

	if btnPressed % 2 == 0 {
		fmt.Println("lampu mati")
	} else {
		fmt.Println("lampu menyala")
	}
}


