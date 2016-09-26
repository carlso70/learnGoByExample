package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		// the more value will be will be false if jobs has been closed
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Recieved job", j)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
