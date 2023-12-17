package a

import (
	"fmt"
)

func f() {
	fmt.Print("test")
	fmt.Printf("test")
	fmt.Println("test") // want "Println is banned!"

	_ = fmt.Sprint("test")
	_ = fmt.Sprintf("test")
	_ = fmt.Sprintln("test")

	Println() // want "Println is banned!"
}

func Println() {}
