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
	f, err := os.Open("./input2.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f);
	counter := 0
	s := make([]int, 3)
	for scanner.Scan() {
		
		line, _ := strconv.Atoi(scanner.Text())
		s = append(s, line)
	}
	
	
	for i := 3; i < len(s)-3; i+=1 {
		totalA := s[i] + s[i+1] + s[i+2]
		totalB := s[i+1] + s[i+2] + s[i+3]
		if totalB > totalA {
			counter++

		}

	}

	fmt.Println(counter)

}
