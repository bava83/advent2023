package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"unicode"
)

func day01a() int{
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

    firstDigit, err := getFirstDigit(s)
    if err != nil{
      panic(err)
    }
    result += firstDigit*10 

    lastDigit, err := getLastDigit(s)
    if err!=nil{
      panic(err)
    }
    result+= lastDigit 
  }

  return result
}

func getFirstDigit(s string) (int, error){
  for _, c := range s{
    if unicode.IsDigit(c){
      result, err := strconv.Atoi(string(c))
      if err != nil{
        return -1, err
      }
      return result, nil
    }
  } 
  return -1, errors.New("Digit not found")
}

func getLastDigit(s string) (int,error){
  for i:=len(s)-1; i>=0; i--{
    c := s[i] 
    if unicode.IsDigit(rune(c)){
      result, err := strconv.Atoi(string(c))
      if err != nil{
        return -1, err
      }
      return result, nil
    }
  } 
  return -1, errors.New("Digit not found")
}
