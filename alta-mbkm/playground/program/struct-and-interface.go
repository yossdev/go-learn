package program

import "fmt"

type Movies struct {
	Title		string
	Director	string
	Year		int
}

type Class struct {
	Title		[]string
}

func MyStructAndInterface()  {
	var avengers *Movies = &Movies{}
	fmt.Println(avengers)
	
	avengers.Title = "Captain america"
	avengers.Director = "Mark Twain"
	avengers.Year = 2009
	fmt.Println(avengers)
	
	
	avengers.Title = "The first avengers"
	avengers.Director = "Russo Brothers"
	avengers.Year = 2012
	fmt.Println(avengers)
	
	loki := &Movies{}
	fmt.Println(loki)
	fmt.Println(*loki)
	
	thor := new(Movies)
	fmt.Println(thor)
	fmt.Println(*thor)
	
	hulk := &Movies{}
	hulk.Title = "the incredible hulk"
	fmt.Println(hulk)
	fmt.Println(&hulk)
	
	var blackwidow Movies
	fmt.Println(blackwidow)
	fmt.Println(&blackwidow)
	
	hawkeye := &Movies{}
	fmt.Println(hawkeye)
	fmt.Println(&hawkeye)
	
	fmt.Println("-----")
}

