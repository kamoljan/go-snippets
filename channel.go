package main

import (
	"fmt"
)

func print(i int) string {
	if i%3 == 0 {
		if i%5 == 0 {
			return fmt.Sprintf("%d PipipiPopopo\n", i)
		}
		return fmt.Sprintf("%d Popopo\n", i)
	} else if i%5 == 0 {
		return fmt.Sprintf("%d Pipip\n", i)
	}
	return fmt.Sprintf("%d\n", i)
}

func loop(f int, t int, c chan string) {
	x := ""
	for i := f; i <= t; i++ {
		x = x + print(i)
	}
	c <- x
}

func main() {
	c := make(chan string)
	go loop(1, 50, c)
	go loop(51, 100, c)
	fmt.Println(<-c, <-c)
}
