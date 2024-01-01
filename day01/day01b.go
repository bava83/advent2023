package main

import (
	"bufio"
	"os"
	"strings"
)

func day01b() int{
  filePath := "./input.txt"
  readFile, err := os.Open(filePath)
  if err != nil {
    panic(err)
  }
  defer readFile.Close()
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  result := 0

  for fileScanner.Scan(){
    s := fileScanner.Text()
    result += getSpellFirstDigit(s)*10 + getSpellLastDigit(s)
  }
  return result
}


var numbers = map[string]int{
  "one":1,"two":2,"three":3,"four":4,"five":5,"six":6,"seven":7,"eight":8,"nine":9,
  "1":1,"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,}

func getSpellFirstDigit(s string) int{
  for i:=0; i<len(s); i++{
    for k,v  := range numbers{
      if strings.HasPrefix(s[i:], k){
        return v 
      }
    } 
  }
  return -1 
}

func getSpellLastDigit(s string) int{
  for i:=len(s)-1; i>=0; i--{
    for k,v := range numbers{
      if strings.HasSuffix(s[:i+1], k){
        return v
      }
    }
  }
  return -1 
}
