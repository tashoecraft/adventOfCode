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
	aim := 0
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		direction := s[0]
		value, _ := strconv.Atoi(s[1])
		if direction == "up" {
			aim -= value
			fmt.Println("up", aim)
		} else if direction == "down" {
			aim += value
			fmt.Println("down", aim)
		} else {
			horizontal += value
			fmt.Println("horizontal", horizontal)
			if aim != 0 {
				vertical += value * aim
			}
			fmt.Println("vertical: ", vertical)

		}
	}

	total := horizontal * vertical

	fmt.Println(total)

}
