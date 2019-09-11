package main

import (
	f "fmt"
	"runtime"
)

func main() {
	var hello = "Hello World"
	f.Println(hello)

	f.Println("Printing Wishes")
	wishes()
	f.Println("\n" + "The number of CPUs is ")
	f.Println(runtime.NumCPU())
}
