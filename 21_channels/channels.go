package main

import (
	"fmt"
	"time"
)

// send
// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("Processing number:", num)
// 		time.Sleep(time.Second)
// 	}
// }

// recived
func revivedChan(result chan int, a int, b int) {
	result <- a * b
}

// process a task
func routineTask(isDone chan bool) {
	defer func() {
		isDone <- true
	}()
	fmt.Println("Task is Running")
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}

	fmt.Println("Task is running complete with sum= ", sum)

}
func task() {

	fmt.Println("Task is Running")
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}

	fmt.Println("Task is running complete with sum = ", sum)

}

// email send fn
func emailSend(emailChan <-chan string, done chan<- bool) {

	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {

	// send message in two channel and recived them
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 100
	}()

	go func() {
		chan2 <- "Eshan"
	}()

	// now print this two channle

	for i := 0; i < 2; i++ {

		select {
		case chan1Val := <-chan1:
			fmt.Println("chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("chan2", chan2Val)
		}
	}

	// emailChan := make(chan string, 100)
	// sendDone := make(chan bool)

	// go emailSend(emailChan, sendDone)

	// for i := 0; i < 3; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("Done Sending all email")
	// // this is important ..................................!
	// close(emailChan)

	// <-sendDone

	// task()
	// task()
	// task()
	// task()

	// isDone := make(chan bool)
	// isDone2 := make(chan bool)
	// go routineTask(isDone)
	// go routineTask(isDone2)
	// go routineTask(make(chan bool))
	// <-isDone
	// <-isDone2

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// recived := make(chan int)

	// go revivedChan(recived, 77849, 999)

	// res := <-recived

	// fmt.Println("recived", <-recived)

}
