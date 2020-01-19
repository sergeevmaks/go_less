package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(50 * time.Millisecond)
	fmt.Println(time.Now())
	delay := 10 * time.Second
	time.Sleep(delay)
	fmt.Println(time.Now())
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}
