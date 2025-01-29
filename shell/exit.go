package shell

import (
	"fmt"
	"os"
)

func Exit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}	