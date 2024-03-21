package main

import (
	"fmt"
	"os"
)

func main() {
	message := os.Getenv("MESSAGE")
	fmt.Println(message)
}
