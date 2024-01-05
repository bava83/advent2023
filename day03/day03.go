package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
    matrix = append(matrix, line)
  }
  return matrix
}

func day03a(matrix [][]rune){
  sum := 0
  for i:=0; i<len(matrix); i++{
    for j:=0; j<len(matrix[i]); j++{
      //found the digit
      if unicode.IsDigit(matrix[i][j]){
        //look for number's length
        numLen := 1
        for k:=j+1; k<len(matrix[i]);k++{
          if unicode.IsDigit(matrix[i][k]){
            numLen++
          }else{
            break
          }
        }

        if checkValid(matrix,i,j,numLen){
          num, err := strconv.Atoi(string(matrix[i][j:j+numLen]))
          if err != nil {
            panic(err)
          }
          sum += num
        }
        //skip the rest of the number
        j += numLen
      }
    }
  }
  fmt.Printf("Day03a sum:%d\n",sum)
}



func checkValid(mt [][]rune, i int, j int, length int)  bool{
  //check left
  if j-1 >= 0 && isValidSymbol(mt[i][j-1]){
     return true
  }
  //check right
  if j+length < len(mt[i]) && isValidSymbol(mt[i][j+length]){
    return true
  }
  //check row above
  if i-1 >= 0{
    //check top left corner
    if j-1 >= 0 && isValidSymbol(mt[i-1][j-1]){
      return true
    } 
    //check top right corner
    if j+length < len(mt[i-1]) && isValidSymbol(mt[i-1][j+length]){
      return true
    }
    //check top center
    for x := 0; x<length; x++{
      if isValidSymbol(mt[i-1][j+x]){
        return true
      }
    }
  }
  //check row below
  if i+1 < len(mt){
    //check bottom left
    if j-1 >= 0 && isValidSymbol(mt[i+1][j-1]){
      return true
    }
    //check bottom right
    if j+length < len(mt[i+1]) && isValidSymbol(mt[i+1][j+length]){
      return true
    }
    //check bottom center
    for x:= 0; x<length; x++{
      if isValidSymbol(mt[i+1][j+x]){
        return true
      }
    }
  }
  return false
}


func isValidSymbol(c rune) bool {
  if unicode.IsGraphic(c) && c != '.' && !unicode.IsDigit(c){
    return true
  }
  return false
}
