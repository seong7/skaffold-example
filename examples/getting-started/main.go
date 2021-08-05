package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello world! 333")

		time.Sleep(time.Second * 5)
	}
}
