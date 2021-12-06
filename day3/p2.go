package main

import (
	"bufio"
	"fmt"
	"os"
        "strings"
        "math"
        "strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type valueTracker struct {
  one []int
  zero []int 
}


func recurse(dataRange []string, indices []int, tieBreak string, bitIndex int) int {
  if len(indices) == 1 {
    return (indices)[0]
  }
  if (len(indices) == 2 && bitIndex > 11) {
    one := strings.Split(dataRange[indices[0]], "")
    
    if tieBreak == "0" {
      if one[11] == "0" {
        return indices[0]
      } else {
        return indices[1]
      } 
    } else {
      if one[11] == "1" {
        return indices[0]
      } else {
        return indices[1]
      }
     }
  }

  tracker := valueTracker{one: []int{}, zero: []int{}}
  
  
  for i:= 0; i < len(indices); i++ {
    index := (indices)[i]
    value := strings.Split(dataRange[index], "")[bitIndex]
    if value == "1" {
      tracker.one = append(tracker.one, index)
    } else {
      tracker.zero = append(tracker.zero, index)
    }
  }

  oneLength := len(tracker.one)
  zeroLength := len(tracker.zero)

  if tieBreak == "1" {
    if oneLength > zeroLength {  
      return recurse(dataRange[:], tracker.one[:], tieBreak, bitIndex+1)
    } else if zeroLength > oneLength { 
      return recurse(dataRange[:], tracker.zero[:], tieBreak, bitIndex+1)
    } else {
      return recurse(dataRange[:], tracker.one[:], tieBreak, bitIndex+1)
    }
  } else {
    if oneLength < zeroLength {  
      return recurse(dataRange[:], tracker.one[:], tieBreak, bitIndex+1)
    } else if zeroLength < oneLength { 
      return recurse(dataRange[:], tracker.zero[:], tieBreak, bitIndex+1)
    } else {
        return recurse(dataRange[:], tracker.zero[:], tieBreak, bitIndex+1)
      }
    }
}

func convertBinary(str string) float64 {
    a := strings.Split(str, "")
      total := float64(0)
    for i := len(str)-1; i >=0; i-- {
      num, _ := strconv.Atoi(a[i])
      if num > 0 {
          pow := float64(len(str)-1-i) 
          total += math.Pow(2, pow) 
      } 
    }
    return total
}


func main() {
	f, err := os.Open("./input1.txt")
	check(err)
	defer f.Close()
	s := bufio.NewScanner(f)
        	
        d := []string{}         
        ii := []int{}
        index := 0
        for s.Scan() {
      	  text := s.Text()
          d = append(d, text)
          ii = append(ii, index)
          index += 1
        }
                ogrIndex := recurse(d, ii, "1", 0)
        coIndex := recurse(d, ii, "0", 0)
        ogrValue := convertBinary(d[ogrIndex])
        coValue := convertBinary(d[coIndex])
        
        fmt.Printf("%f\n", ogrValue * coValue)
  }


