package single

import (
	"fmt"
)

func as() {
	fmt.Print("test")
	fmt.Printf("test")
	fmt.Println("test") // want "Println is banned!"

	_ = fmt.Sprint("test")
	_ = fmt.Sprintf("test")
	_ = fmt.Sprintln("test")
}
