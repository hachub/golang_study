package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func controller(control chan<- string, results <-chan string) {
	for j := 0; j < 9; j++ {

		fmt.Printf("\nSTART WORKERS: \n")
		for i := 0; i < 10; i++ {
			fmt.Printf("[%d] send next \n", i)
			control <- "next"
		}

		fmt.Printf("START LISTEN \n")
		for i := 0; i < 10; i++ {
			msg := <-results
			if i == 0 {
				fmt.Printf("CONTROLLER RES: ")
			}
			fmt.Printf("[%d] - %s ", i, msg)
		}

	}
	fmt.Printf("\nEND PROCESS \n")
}

func worker(id int, control <-chan string, results chan<- string) {
	for i := 0; i < 10; i++ {
		// wait control signal
		msg := <-control
		fmt.Printf("[%d] worker %d - get msg %s \n", i, id, msg)

		// DO WORK
		td := time.Duration((rand.Intn(250)))
		time.Sleep(time.Millisecond * td)

		// send result
		results <- strconv.Itoa(id)

	}
}
func multithreading() {

	fmt.Println("-------------- Multithreading -------------------")

	var results chan string = make(chan string)
	var control chan string = make(chan string)

	for i := 0; i < 10; i++ {
		go worker(i, control, results)
	}
	go controller(control, results)

	var input string
	fmt.Scanln(&input)
}
