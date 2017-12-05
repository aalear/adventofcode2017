package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func panicOnError(e error) {
	if(e != nil) {
		panic(e)
	}
}

func main() {
	f, err := ioutil.ReadFile("input.txt")
	panicOnError(err)
	
	input := string(f)
	sum := 0
	
	first, err := strconv.Atoi(input[:1])
	panicOnError(err)

	for i := 1; i < len(input); i++ {
		curr, err := strconv.Atoi(input[i - 1:i])
		panicOnError(err)

		next, err := strconv.Atoi(input[i:i + 1])
		panicOnError(err)

		if curr == next {
			sum += curr
		}
	}

	last, err := strconv.Atoi(input[len(input) - 1:])
	if last == first {
		sum += last
	}

	fmt.Println(sum)
}