package main

import (
	"fmt"
	"time"
)

func using_select() {
	fmt.Println("-------------- Multithreading select -------------------")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second * 5)
			c1 <- ""
		}
	}()

	go func() {
		for {
			time.Sleep(time.Minute)
			c2 <- "the minute is up"
		}
	}()

	c := 5
	for {
		select {
		case <-c1:
			if c > 0 {
				fmt.Printf("%d", c)
			}
			c += 5
		case msg := <-c2:
			fmt.Printf("%d %s \n", c, msg)
			c = 0
		case <-time.After(time.Second):
			fmt.Print(".")
		}
	}

	var input string
	fmt.Scanln(&input)

}
