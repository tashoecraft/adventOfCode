package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./input1.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	vertical := 0
	horizontal := 0
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		direction := s[0]
		value, _ := strconv.Atoi(s[1])
		if direction == "up" {
			vertical -= value
		} else if direction == "down" {
			vertical += value
		} else {
			horizontal += value

		}
	}

	total := horizontal * vertical

	fmt.Println(total)

}
