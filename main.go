package main

import (
	"encoding/json"
	"strings"
	"log"
	"os"
	"strconv"
	"time"
	"bufio"
	"fmt"
)

type Answer struct {
	Sequence string `json:"Sequence of the number"`
	Number string  `json:"Fibonacci number"`
}


func main() {

	
	fibonacciList := map[int]int{1: 0, 2: 1}
	correctCounter := 0
	mistakeCounter := 0
	sequence := 0
	c := make(chan string, 1)


	nextFibonacci(fibonacciList[1], fibonacciList[2], fibonacciList, mistakeCounter, sequence, correctCounter, c)
}

func nextFibonacci(f1 int, f2 int, nums map[int]int, correctCounter int, mistakeCounter int, sequence int, c chan string) func() {
	if correctCounter == 10 {
		fmt.Println("10 correct answers - you won!")
		os.Exit(1)
	} else if mistakeCounter == 3 {
		fmt.Println("3 mistakes - you lost")
		os.Exit(1)
	}

	sequence++
	
	ans := Answer{strconv.Itoa(sequence), strconv.Itoa(nums[sequence])}
		a, err := json.Marshal(ans)
	if err != nil {
		log.Fatal(err)
	}
		
	fmt.Println("Correct answers: ", correctCounter)
	fmt.Println("Mistakes: ", mistakeCounter)
	
	f3 := f1 + f2
	n := len(nums) + 1
	nums[n] = f3

	go input(c)

	fmt.Println("You have 10 seconds to answer")

	for {
		select {
		case i := <-c:
			if strconv.Itoa(nums[sequence]) == i {
				fmt.Println("Correct answer")
				fmt.Println()
				correctCounter++
				return nextFibonacci(nums[n-1], nums[n], nums, correctCounter, mistakeCounter, sequence, c)
			} 
				
				fmt.Println("Wrong answer. Next number is: ")
				fmt.Printf("%s\n", a)
				fmt.Println()
				mistakeCounter++
				return nextFibonacci(nums[n-1], nums[n], nums, correctCounter, mistakeCounter, sequence, c)
							

		case <-time.After(10000 * time.Millisecond):
			mistakeCounter++
			fmt.Print("\nTime out. Next number is: ")
			fmt.Printf("%s\n", a)
			fmt.Println()
			return nextFibonacci(nums[n-1], nums[n], nums, correctCounter, mistakeCounter, sequence, c)
			}
			
		
	}
}

func input(c chan string) {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Enter next Fibonacci number: ")
	text, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	text = strings.Trim(text," \r\n")
	c <- text
}
