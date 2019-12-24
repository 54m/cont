package main

import (
	"fmt"
	"github.com/54mch4n/cont"
)

func main() {
	var (
		ex1 = []string{"1", "2", "3"}
		ex2 = []int{1, 2, 3}
		ex3 = "123"
	)
	fmt.Println(ex1, ex2, ex3)
	if cont.Contains(ex1, "1") {
		fmt.Println("ex1: OK!")
	}
	if cont.Contains(ex2, 1) {
		fmt.Println("ex2: OK!")
	}
	if cont.Contains(ex3, "1") {
		fmt.Println("ex3: OK!")
	}
}
