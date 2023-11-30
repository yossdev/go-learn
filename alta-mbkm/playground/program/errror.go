package program

import (
	"errors"
	"fmt"
)

func Error() {
	var err error = errors.New("ada error")
	if err == nil {
		fmt.Println(err)
	} else {
		fmt.Println("aman")
	}
}