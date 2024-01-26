package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	sum := 0
	dat, err := os.Open("../example.txt")
	check(err)
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		var numbers []int
		lineNum := 0
		for _, s := range scanner.Text() {
			if dig, err := strconv.Atoi(string(s)); unicode.IsNumber(s) {
				check(err)
				fmt.Printf("converted digit: %d string: %s \n", dig, string(s))
				numbers = append(numbers, dig)
			}
		}
		if len(numbers) > 1 {
			lineNum, err = strconv.Atoi(fmt.Sprintf("%d%d", numbers[0], numbers[len(numbers)-1]))
			check(err)
			sum = sum + lineNum
		} else if len(numbers) == 1 {
			lineNum, err = strconv.Atoi(fmt.Sprintf("%d%d", numbers[0], numbers[0]))
			check(err)
			sum = sum + lineNum
		}

		fmt.Printf("line from file: %s calib: %d \n", scanner.Text(), lineNum)
	}

	fmt.Printf("Sum: %d \n", sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
