package main

import (
	"log"
	"time"

	"github.com/54mch4n/cont"
)

func main() {
	var (
		now = time.Now()
		ex1 = []string{"1", "2", "3"}
		ex2 = []int{1, 2, 3}
		ex3 = "123"
		ex4 = map[string]interface{}{"a": "1", "b": "2", "c": "3", "d": now}
	)

	log.Println(ex1)
	log.Println(ex2)
	log.Println(ex3)
	log.Println(ex4)

	if cont.Contains(ex1, "1") {
		log.Println("ex1: OK!")
	}

	if cont.Contains(ex2, 1) {
		log.Println("ex2: OK!")
	}

	if cont.Contains(ex3, "1") {
		log.Println("ex3: OK!")
	}

	if cont.Contains(ex4, now) {
		log.Println("ex4: OK!")
	}
}
