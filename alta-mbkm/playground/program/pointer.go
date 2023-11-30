package program

import "fmt"

func MyPointer()  {
	//* init pointer
	// var n int = 5
	// var addressN *int = &n

	// fmt.Println(n)
	// // fmt.Println(&n)
	// fmt.Println(*addressN)


	// var x int = 10
	// var addressX *int = &x

	// fmt.Scanln(&x)

	// fmt.Println(x)
	// fmt.Println(&x)
	// fmt.Println(*addressX)

	// addressN = &x
	// addressX = &n
	// fmt.Println(*addressN)
	// fmt.Println(*addressX)

	var val1,val2,temp int32

	fmt.Print("\n")
	fmt.Print("\tSwap Two Numbers Using Pointers")
	fmt.Print("\n\n")
	fmt.Print("\tEnter first value   : ")
	fmt.Scanf("%d\n",&val1)
	fmt.Print("\tEnter second value  : ")
	fmt.Scanf("%d\n",&val2)
	fmt.Print("\n")
	fmt.Print("\t===== BEFORE SWAPPING ===== ")
	fmt.Print("\n\n")
	fmt.Print("\tA = ",val1," and B = ",val2)
	fmt.Print("\n\n")

	a := &val1;
	b := &val2;

	temp = *a;
	*a = *b;
	*b = temp;

	fmt.Print("\t===== AFTER SWAPPING ===== ")
	fmt.Print("\n\n")
	fmt.Print("\tA = ",val1," and B = ",val2)
	fmt.Print("\n\n")
	fmt.Print("\tEnd of Program")
	fmt.Print("\n")

}