package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("it's Windows!")
	case "darwin":
		fmt.Println("it's Mac OS!")
	case "linux":
		fmt.Println("it's Linux!")
	default:
		fmt.Printf("%s.\n", os)
	}
}
