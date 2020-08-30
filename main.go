package main

import (
	"fmt"
	"go-face-reko-aws/utils"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(utils.CompareFaces("go-face-reko-aws", "1.png", "2.png"))
}
