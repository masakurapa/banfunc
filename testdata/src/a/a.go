package a

import (
	"fmt"
)

func as() {
	fmt.Print("test")
	fmt.Printf("test")
	fmt.Println("test") // want "Println is banned!"

	_ = fmt.Sprint("test")
	_ = fmt.Sprintf("test") // want "Sprintf is banned!"
	_ = fmt.Sprintln("test")
}
