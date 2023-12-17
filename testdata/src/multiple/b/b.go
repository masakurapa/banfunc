package b

import (
	"fmt"
)

func f() {
	fmt.Print("test")
	fmt.Printf("test")
	fmt.Println("test") // want "fmt.Println is banned!"

	_ = fmt.Sprint("test")
	_ = fmt.Sprintf("test") // want "fmt.Sprintf is banned!"
	_ = fmt.Sprintln("test")

	Println()
}

func Println() {}
