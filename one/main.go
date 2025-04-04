package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []int {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(input)
	inputs := make([]int, 5, 5)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if line == "" {
			break
		}

		number, e := strconv.Atoi(line)
		if e != nil {
			panic(e)
		}

		inputs = append(inputs, number)
		if err != nil {
			break
		}
	}
	defer input.Close()

	return inputs
}

func partOne() {
	fmt.Println("aoc #1a")

	var inputs []int = readInput()

	var sum int64
	for i := range inputs {
		sum += int64(inputs[i])
	}

	fmt.Println(sum)
}

func main() {
	partOne()
}
