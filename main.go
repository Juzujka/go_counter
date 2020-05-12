package main

import (
	"fmt"
	"os"

	"go_counter/go_counter_proc"
)

func main() {
	fmt.Printf("Total: %d\n", go_counter_proc.InputCountGo(os.Stdin, 5))
}
