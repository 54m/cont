package main

import (
	"fmt"
	"time"

	"github.com/54mch4n/cont"
)

func main() {
	ExampleContainsStringSlice()
	ExampleContainsIntSlice()
	ExampleContainsString()
	ExampleContainsMap()
}

func ExampleContainsStringSlice() {
	contents := []string{"1", "2", "3"}
	fmt.Printf("%#v\n", contents)
	if cont.Contains(contents, "1") {
		fmt.Println("StringSlice: OK!")
	}
	// Output: []string{"1", "2", "3"}
	// Output: StringSlice: OK!
}

func ExampleContainsIntSlice() {
	contents := []int{1, 2, 3}
	fmt.Printf("%#v\n", contents)
	if cont.Contains(contents, 1) {
		fmt.Println("IntSlice: OK!")
	}
	// Output: []int{1, 2, 3}
	// Output: IntSlice: OK!
}

func ExampleContainsString() {
	contents := "123"
	fmt.Printf("%#v\n", contents)
	if cont.Contains(contents, "1") {
		fmt.Println("String: OK!")
	}
	// Output: "123"
	// Output: String: OK!
}

func ExampleContainsMap() {
	now := time.Now()
	contents := map[string]interface{}{"a": "1", "b": 2, "c": "3", "d": now}
	fmt.Printf("%#v\n", contents)
	if cont.Contains(contents, now) {
		fmt.Println("Map: OK! - time.Time")
	}
	if cont.Contains(contents, 2) {
		fmt.Println("Map: OK! - int")
	}
	if cont.Contains(contents, "c") {
		fmt.Println("Map: OK! - string")
	}
	// nolint:lll
	// Output: map[string]interface {}{"a":"1", "b":2, "c":"3", "d":time.Time{wall:0x1234567890, ext:xxx, loc:(*time.Location)(0x123456789)}}
	// Output: Map: OK! - time.Time
	// Output: Map: OK! - int
	// Output: Map: OK! - string
}
