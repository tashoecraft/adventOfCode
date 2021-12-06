package main

import (
	"bufio"
	"fmt"
	"os"
        "strings"
        "math"
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
	plus := [12]int{0,0,0,0,0,0,0,0,0,0,0,0}
	for scanner.Scan() {
		s := scanner.Text()
                a := strings.Split(s, "")

                for i, r := range a {
                  if r == "1" {
                    plus[i]++
                  } else {
                    plus[i]--
                  }
                }
	}
        gamma := float64(0)
        epsilon := float64(0)
        for i := len(plus)-1; i >=0; i-- {
          if plus[i] > 0 {
            pow := float64(len(plus)-1-i) 
            gamma += math.Pow(2, pow) 
          } else {
            pow := float64(len(plus)-1-i) 
            epsilon += math.Pow(2, pow) 

          }
        }
        fmt.Printf("%f\n", gamma * epsilon)
}
