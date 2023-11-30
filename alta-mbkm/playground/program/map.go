package program

import "fmt"

func BelajarMap() {
	//* Deklarasi Map isi manual
	var myMap = map[int]string{1: "id", 2: "nama"}
	myMap[3] = "jurusan"
	fmt.Println(myMap)
	fmt.Println(myMap[3])

	//* Deklarasi Map isi dinamis using make keyword
	var myMap2 = make(map[int]string)
	fmt.Println(myMap2)
	myMap2[1] = "id"
	fmt.Println(myMap2)

	//* mengisi map dari loop
	var i int
	myMap3 := map[int]int{}
	for i < 10 {
		myMap3[i] = i + 1
		i++
	}
	fmt.Println(myMap3)

	//* looping untuk mencari apakah suatu bilangan ada di dalam map
	var j int
	for j < len(myMap3) {
		if myMap3[j] == 8 {
			fmt.Println("Ada")
		} else {
			fmt.Println("Tidak Ada")
		}
		j++
	}
	
}