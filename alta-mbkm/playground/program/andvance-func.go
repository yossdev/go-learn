package program

import "fmt"

func AdvanceFunc() {

	resVariadic := variadic(2, 4, 5)
	fmt.Println(resVariadic)
// 	valueAnonim := anonim()
// 	fmt.Println(value())
}

//* variadic function
func variadic(number ...int) int {
	var total int
	for _, number := range number {
		total += number
	}
	return total
}


//* anonymous/literal function
func anonim() {
	func () {
		fmt.Println("Ini fungsi anonim")
	} ()
}


//* closure (provide data isolation)

//* defer (calls when the function is finish)