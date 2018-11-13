package main

import (
	"fmt"
)

func helloDC(done chan bool) {
	fmt.Println("Hello Hearst DC Offsite!!!!!!!!!!!!!!")
	done <- true
}
func main() {
	done := make(chan bool)
	go helloDC(done)
	<-done
	fmt.Println("main function")
}
