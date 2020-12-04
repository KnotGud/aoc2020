package main

import (
	"fmt"
	"strconv"
)

func part1(in <-chan string) {
	c := make(chan int, 10)
	go func() {
		for s := range in {
			if n, err := strconv.Atoi(s); err == nil {
				c <- n
			}
		}

		close(c)
	}()

	data := make(map[int]bool)
	for num := range c {
		data[num] = true
	}

	for i := range data {
		temp := 2020 - i
		if data[temp] {
			fmt.Printf("Values %v * %v = %v\n", i, temp, i*temp)
			break
		}
	}
}
