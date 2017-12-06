package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"strings"
	"math"
	"strconv"
)

func panicOnError(e error) {
	if(e != nil) {
		panic(e)
	}
}

func main() {
	f, err := os.Open("input.txt")
	panicOnError(err)
	
	input := bufio.NewReader(f)

	checksum := 0
	for {
		row, err := input.ReadString('\n');
		if err != nil && err != io.EOF {
			panicOnError(err)
		}
		
		rowArray := strings.Split(row, "\t")

		min := math.MaxInt32
		max := math.MinInt32
		for i := 0; i < len(rowArray); i++ {
			num, err := strconv.Atoi(strings.Trim(rowArray[i], "\r\n"))
			panicOnError(err)

			if(num < min) {
				min = num
			}
			if(num > max) {
				max = num
			}
		}

		checksum += (max - min)

		if err == io.EOF { 
			break
		}
	}

	fmt.Println(checksum)
}