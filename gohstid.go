package main

import (
	"fmt"
	"time"
)

func main() {
	v := 1
	for {
		time.Sleep(1111 * time.Millisecond)
		v += 1
		fmt.Println(v)
	}
}
