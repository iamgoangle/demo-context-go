package main

import (
	"fmt"
)

func main() {
	var i int
	go func() {
		i = 5
	}()

	fmt.Println(i)
}
