package main


import (
	"bufio"
	"os"
	"fmt"
	"strconv"
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
	scanner := bufio.NewScanner(f);
	previousValue := 0 
	counter := 0

	for scanner.Scan() {
		
		line, _ := strconv.Atoi(scanner.Text())
		if line > previousValue {
			
			fmt.Println(line, previousValue, counter)
			counter++
		}
		previousValue = line
	}
	fmt.Println(counter)

}
