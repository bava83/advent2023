package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strconv"
	"unicode"
)

func day03(){
  filePath := "input.txt"
  matrix := readFilesToSlice(filePath)
  day03a(matrix)
}


func readFilesToSlice(filePath string) [][]rune{
  readFile, err := os.Open(filePath)
  if err != nil {
    panic(err)
  }
  defer readFile.Close()

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines) 

  var matrix [][]rune
  for fileScanner.Scan(){
    l := fileScanner.Text()
    var line []rune
    for _,c := range l {
      line = append(line, c)
    } 
    fmt.Println(line)
    matrix = append(matrix, line)
  }
  return matrix
}

func day03a(matrix [][]rune){
  sum := 0
  for i:=0; i<len(matrix); i++{
    for j:=0; j<len(matrix[i]); j++{
      if unicode.IsDigit(matrix[i][j]){
        numLen := 1
        for k:=j+1; k<len(matrix[i]);k++{
          
        }
      }
    }
  }
}
