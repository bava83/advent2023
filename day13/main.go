package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
  //filePath := "./input.txt"
  filePath := "./input02.txt"
  m := readData(filePath)
  day13a(m)
}

func readData(filepath string)[][]string{
  readFile, err := os.Open(filepath)
  if err != nil{
    panic(err)
  }
  defer readFile.Close()

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  var matrix [][]string
  var line []string
  for fileScanner.Scan(){
    l := fileScanner.Text()
    if l != ""{
      line = append(line, l)
    }else {
      matrix = append(matrix, line)
      line = []string{}
    }
  }
  matrix = append(matrix, line)
  return matrix
}

func findReflection(m [][]rune)int{
  result := 0
//  v := 0
  h, foundH := findPosition(m)
  if !foundH{
    fmt.Println("didn't found horizontal mirror")
    //transpost the matrix, 
//    v := findVertical(m)
  }
  fmt.Println(h)
  return result
}

func findPosition(m [][]rune)(int,bool){
  result := 0
  for i:= 0; i<len(m)-1; i++{
    if checkEqual(m[i],m[i+1]){
      flag := true
      for j,k:=i,i+1; j>0 && k<len(m)-1 ; j,k = j-1,k+1{
        if(!checkEqual(m[j],m[k])){
          flag = false
          break
        }
      }
      if flag == true{
        return i+1, true
      }
    }
  }
  return result, false
}

func checkEqual(a []rune,b []rune)bool{
  for i := range a{
    if a[i] != b[i]{
      return false
    }
  }
  return true
}

func convert2Rune(m []string)[][]rune{
  var i [][]rune
  var j []rune
  for _, row := range m{
    for _, col := range row{
      j = append(j, col)
    }
    i = append(i, j)
    j = []rune{}
  }
  return i
}

func day13a(m [][]string){
  ru := convert2Rune(m[1])
  r := findReflection(ru)
  fmt.Println(r)
}
