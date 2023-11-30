package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	//* User Input
	// fmt.Println(Compare("", ""))
	// fmt.Println(caesar(0, ""))

	//! Test Hardcoded
	//* Problem 1
	// fmt.Println(Compare("AKA", "AKASHI"))

	//* Problem 2
	// fmt.Println(caesar(3, "abc"))
	// fmt.Println(caesar(2, "alta"))
	// fmt.Println(caesar(10, "alterraacademy"))
	// fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	// fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

	//* Problem 3
	// a := 10
	// b := 20
	// swap(&a, &b)
	// fmt.Println(a, b)

	//*Problem 4
	// var a1, a2, a3, a4, a5, a6, min, max int
	// fmt.Scan(&a1)
	// fmt.Scan(&a2)
	// fmt.Scan(&a3)
	// fmt.Scan(&a4)
	// fmt.Scan(&a5)
	// fmt.Scan(&a6)
	// min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	// fmt.Println("Nilai min ", min)
	// fmt.Println("Nilai max ", max)

	//* Problem 5
	// var a = Student{}
	// for i := 0; i < 5; i++ {
	// 	var name string
	// 	fmt.Print("Input " + string(i) + " Student's Name : ")
	// 	fmt.Scan(&name)
	// 	a.name = append(a.name, name)
	// 	var score int
	// 	fmt.Print("Input " + name + " Score : ")
	// 	fmt.Scan(&score)
	// 	a.score = append(a.score, score)
	// }
	// fmt.Println(a)
	// fmt.Println("\n\nAverage Score Students is", a.Avarage())
	// fmt.Println("\n\nAverage Score Students is", Avarage(a))
	// scoreMax, nameMax := a.Max()
	// fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	// scoreMin, nameMin := a.Min()
	// fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")

	//* Problem 6
	// var menu int
	var a = student{}
	fmt.Println(a)
	a.name = "tes"
	fmt.Println(a)
	var c Chiper = &a
	fmt.Println(c)
	// fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	// fmt.Scan(&menu)
	// if menu == 1 {
	// 	fmt.Print("\nInput Student's Name : ")
	// 	fmt.Scan(&a.name)
	// 	fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	// } else if menu == 2 {
	// 	fmt.Print("\nInput Student's Encode Name : ")
	// 	fmt.Scan(&a.nameEncode)
	// 	fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	// } else {
	// 	fmt.Println("Wrong input name menu")
	// }
}

//* Problem 1 - Compare String
/*
	@param a dan b tidak menerima angka
	@param menerima lowercase dan uppercase
	@param dapat berupa mixed string
	@proses case sensitive
*/
func Compare(a, b string) string {
	defer func()  {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Masukkan String: ")
	fmt.Scanf("%s\n", &a)
	_, errA := strconv.Atoi(string(a))
	if errA == nil {
		panic("Tidak boleh diisi angka, coba lagi...") //! This will shutdown the program
	}

	fmt.Println("Masukkan String: ")
	fmt.Scanf("%s\n", &b)
	_, errB := strconv.Atoi(string(b))
	if errB == nil {
		panic("Tidak boleh diisi angka, coba lagi...") //! This will shutdown the program
	}

	if len(a) == 0 || len(b) == 0 {
		return "No common Substring"
	}

	if strings.Contains(a, b) || strings.Contains(b, a) {
		if len(a) > len(b) {
			return b
		} else {
			return a
		}
	}
	return "No common Substring"
}


//* Problem 2 - Caesar Cipher
func caesar(offset int, input string) string {
	defer func()  {
		if err := recover(); err != nil {
			fmt.Println("Bukan merupakan angka, coba lagi...")
		}
	}()

	fmt.Println("Masukkan nilai offset: ")
	fmt.Scanf("%d\n", &offset)

	fmt.Println("Masukkan String: ")
	fmt.Scanf("%s\n", &input)
	_, errB := strconv.Atoi(input)
	if errB == nil {
		return "Bukan merupakan string, coba lagi..."
	}
	
	dict := make([]int, 0)
	for i := 0; i < len(input); i++ {
		dict = append(dict, int(input[i]))
	}

	start := offset % len(input)

	if start != 0 {
		for k := 0; k < start; k++ {
			for m := 0; m < offset; m++ {
				dict[k] += 1
				if dict[k] == 123 {
					dict[k] = 97
				}
			}
			fmt.Printf("%s", string(dict[k]))
		}
	}
	for j := start; j < len(dict); j++ {
		for m := 0; m < offset; m++ {
			dict[j] += 1
			if dict[j] == 123 {
				dict[j] = 97
			}
		}
		fmt.Printf("%s", string(dict[j]))
	}
	return ""
}


//* Problem 3 - Swap Two Number Using Pointer
func swap(a, b *int) {
	*a, *b = *b, *a
}


//* Problem 4 - Min and Max Using Pointer
func getMinMax(numbers ...*int) (min int, max int) {
	resMin := *numbers[0]
	for i, _ := range numbers {
		if *numbers[i] < resMin {
			resMin = *numbers[i]
		}
	}

	resMax := *numbers[0]
	for i, _ := range numbers {
		if *numbers[i] > resMax {
			resMax = *numbers[i]
		}
	}
	return resMin, resMax
}


//* Problem 5 - Students Score
type Student struct {
	name 	[]string
	score 	[]int
}

// method style
func (s Student) Avarage() float64 {
	sumMap := make(map[int]float64)
	var sum int
	for i, value := range s.score {
		sumMap[i] =  float64(sum) + float64(value)
		sum += value
	}
	return sumMap[len(sumMap)-1]/float64(len(s.score))
}

// function style
func Avarage(s Student) float64 {
	sumMap := make(map[int]float64)
	var sum int
	for i, value := range s.score {
		sumMap[i] =  float64(sum) + float64(value)
		sum += value
	}
	return sumMap[len(sumMap)-1]/float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	scoreMap := make(map[int]int)
	minStudent := 0
	minScore := s.score[0]
	for i, value := range s.score {
		scoreMap[i] = + value
		if minScore > value {
			minScore = scoreMap[i]
			minStudent = i
		}
	}
	return minScore, s.name[minStudent]
}

func (s Student) Max() (max int, name string) {
	scoreMap := make(map[int]int)
	maxStudent := 0
	maxScore := s.score[0]
	for i, value := range s.score {
		scoreMap[i] = value
		if maxScore < value {
			maxScore = scoreMap[i]
			maxStudent = i
		}
	}
	return maxScore, s.name[maxStudent]
}


//! Pelajari lagi INI!
//* Problem 6 - Substitution Cipher
//* @param hanya menerima lowercase
type student struct {
	name 		string
	nameEncode 	string
	score 		int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	// code here
	dict := make([]int, 0)
	for i := 0; i < len(s.name); i++ {
		dict = append(dict, int(s.name[i]))
	}

	offset := 17
	start := offset % len(s.name)

	if start != 0 {
		for k := 0; k < start; k++ {
			for m := 0; m < offset; m++ {
				dict[k] += 1
				if dict[k] == 123 {
					dict[k] = 97
				}
			}
			nameEncode += string(dict[k])
		}
	}
	for j := start; j < len(dict); j++ {
		for m := 0; m < offset; m++ {
			dict[j] += 1
			if dict[j] == 123 {
				dict[j] = 97
			}
		}
		nameEncode += string(dict[j])
		
	}
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	dict := make([]int, 0)
	for i := 0; i < len(s.nameEncode); i++ {
		dict = append(dict, int(s.nameEncode[i]))
	}

	offset := 9
	start := offset % len(s.nameEncode)

	if start != 0 {
		for k := 0; k < start; k++ {
			for m := 0; m < offset; m++ {
				dict[k] += 1
				if dict[k] == 123 {
					dict[k] = 97
				}
			}
			nameDecode += string(dict[k])
		}
	}
	for j := start; j < len(dict); j++ {
		for m := 0; m < offset; m++ {
			dict[j] += 1
			if dict[j] == 123 {
				dict[j] = 97
			}
		}
		nameDecode += string(dict[j])
	}
	
	return nameDecode
}

