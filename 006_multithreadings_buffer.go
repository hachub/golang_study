package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func qworker(id int, tasks chan string) {
	for {
		msg := <-tasks
		duration := time.Second * time.Duration(rand.Intn(10))
		time.Sleep(duration)
		fmt.Printf("worker %d end %s Duration %v sec\n", id, msg, duration)
	}
}

func using_buffer() {
	fmt.Println("-------------- Multithreading buffer -------------------")

	tasks := make(chan string, 10)

	// start workers
	for i := 1; i < 3; i++ {
		go qworker(i, tasks)
	}

	taskNum := 1

	for {
		task := "task" + strconv.Itoa(taskNum)
		taskNum++
		time.Sleep(time.Second * time.Duration(rand.Intn(2)))
		tasks <- task
		fmt.Printf(" -> tasks: %d [ %d ]\n", len(tasks), cap(tasks))

	}

	var input int
	fmt.Scanln(&input)
}
