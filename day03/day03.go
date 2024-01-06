package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func day03(){
  filePath := "input.txt"
  matrix := readFilesToSlice(filePath)
  a := day03a(matrix)
  fmt.Printf("Day03a sum:%d\n",a)
  b := day03b(matrix)
  fmt.Printf("Day03b product:%d\n",b)
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

func day03a(matrix [][]rune)int{
  sum := 0
  for i:=0; i<len(matrix); i++{
    for j:=0; j<len(matrix[i]); j++{
      //found the digit
      if unicode.IsDigit(matrix[i][j]){
        //look for number's length
        numLen := findLength(matrix,i,j,1)

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
  return sum
}

func checkValid(mt [][]rune, i int, j int, length int) bool{
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

func day03b(mt [][]rune)int{
  result := 0
  //find *
  for i:=0; i< len(mt); i++{
    for j:=0; j< len(mt[i]); j++{
      if mt[i][j] == '*' {
        //look for the pair 
        result += findPair(mt,i,j)
      }
    }
  }
  return result
}

type Point struct{
  x int
  y int
}

//find the pair of number to multiply
//return 0 when didn't find the pair
//return the multiplied value when found
func findPair(mt [][]rune, i, j int) int{
  const PAIR = 2 
  result := 1
  indices := []Point{
    {-1,-1},{-1,0},{-1,1},
    {0,-1},/*symbol*/{0,1},
    {1, -1}, {1, 0}, {1, 1},
  }

  //store the starting point of a number, and the end position
  num := make(map[Point]int)

  for _, idx := range indices{
    checkRow := i+idx.x
    checkCol := j+idx.y

    //if the checking position is within boundary
    if checkRow >= 0 && checkRow < len(mt) && checkCol >=0 && checkCol < len(mt[i]){
      if unicode.IsDigit(mt[checkRow][checkCol]){
        start, end := findNumPosition(mt,checkRow,checkCol)
        p := Point{checkRow,start}
        num[p] = end
      }
    }
  }

  if len(num)!=PAIR{
    return 0
  }

  //found a pair
  for k,v := range num{
    x,err := strconv.Atoi(string(mt[k.x][k.y:v])) 
    if err != nil{
      panic(err)
    }
    result *= x
  }

  return result
}



//given a number's position, return it's starting position and ending position
func findNumPosition(mt [][]rune, i, j int) (int, int){
  length := 1
  //find digit before current position
  for x:=j-1; x>=0; x--{
    if !unicode.IsDigit(mt[i][x]){
      break
    }   
    length++
  }
  start := j-length+1
  
  //find digit after current position
  length = findLength(mt,i,j,length)

  return start,start+length
}

//given the previous lenght, we want to find the length of the rest of number
func findLength(mt [][]rune, i,j,length int) int{
  for k:=j+1; k<len(mt[i]);k++{
    if !unicode.IsDigit(mt[i][k]){
      break
    }
    length++
  }
  return length
}
