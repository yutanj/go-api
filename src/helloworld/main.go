package main

import (
	"fmt"

	"github.com/yutanj/go_practice/calculator"
)

func main() {
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)
}
