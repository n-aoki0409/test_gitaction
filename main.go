package main

import (
	"fmt"
	"os"
)

func main() {
	var stage string
	message := os.Getenv("MESSAGE")
	flag.StringVar(&stage, "stage", "", "stage name")
	fmt.Println(stage)
	fmt.Println(message)
}
