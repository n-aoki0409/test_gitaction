package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var stage string
	message := os.Getenv("MESSAGE")
	flag.StringVar(&stage, "stage", "", "stage name")
	flag.Parse()
	fmt.Println(stage)
	fmt.Println(message)
}
