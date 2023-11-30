package main

func main() {
	// fmt.Println(b.PrimeNum())

	// b.LampBtn()

	// b.BelajarMap()

	// b.PlayWithString()

	// b.AdvanceFunc()

	// b.MyPointer()

	// b.MyMethod()

	// b.BitwiseOperator()

	// b.Error()

	// program.MyStructAndInterface()

	// slice := make([]int, 0)
	// slice = append(slice, 1, 5)
	// fmt.Println(cap(slice))
	// slice[0] = slice[1]
	// slice[1] = slice[0]
	// fmt.Println(slice)
	// fmt.Println(len(slice))

	//* Another way assign value to array or slice
	// arr := [5]int{1,2,3,4,5}
	// fmt.Println(arr)
	// arr[0] = arr[3]
	// arr[3] = arr[0]
	// fmt.Println(arr)
}

// func main() {
//     res := return()
//     x, y := swap(3, 4)
//     fmt.Println("start")
//     defer welcome()
//     for i := 0; i < 5; i++ {
//         defer fmt.Println(i)
//     }
//     fmt.Println("End")
//     set()
//     fmt.Println(x)

//     pointer()

//     fmt.Println(primeNum())
// }

//* Learning about pointer
// func pointer()  {
// 	var p *int
// 	x := 42
//     x += 1000
// 	p = &x
// 	fmt.Println(p)
//     fmt.Println(*p)

//* changing x value from pointer using * (dereferencing operator)
//     *p = 8
//     fmt.Println(p)
//     fmt.Println(*p)
//     fmt.Println(x)
// }

// func return() int {
//     var num int
//     fmt.Scanln(&num)
//     return num
// }

// func swap(x, y int) (int, int) {
//     return y, x
// }

// func welcome() {
//     fmt.Println("Main finished")
// }

// var x int = 2
// func set() {
//     x += 1
// }

//* Pola Bintang Siku-siku
// func Bintang() {
// 	var str string
// 	for i := 0; i < 5; i++ {
// 		for j := 0; j < i; j++ {
// 			str += "*"
// 		}
// 		fmt.Println(str)
// 		str = ""
// 	}
// }