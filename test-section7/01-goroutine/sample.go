package main

import "fmt"

// func goroutine(s []string, c chan int) {
func goroutine(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}

	close(c)
}

/**
 * 以下を出力させる
 * test1!
 * test1!test2!
 * test1!test2!test3!
 * test1!test2!test3!test4!
 **/
func main() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	// c := make(chan int)
	c := make(chan string)
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}
}
