package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
  filePath := "./input.txt"
  //filePath := "./input02.txt"
  m := readData(filePath)
  day13a(m)
  day13b(m)
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
  h, foundH := findMirrorPosition(m)
  result += h*100
  if !foundH{
    tm := transpost(m)
    v, foundV := findMirrorPosition(tm)
    if !foundV{
      panic("found neither")
    }
    result += v
  }
  return result
}

func findMirrorPosition(m [][]rune)(int,bool){
  for i:= 0; i<len(m)-1; i++{
    if checkEqual(m[i],m[i+1]){
      flag := true
      //2 pointers, moving towards either end
      for j,k:=i,i+1; j>=0 && k<len(m) ; j,k = j-1,k+1{
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
  return 0, false
}

func checkEqual(a []rune,b []rune)bool{
  return strings.Compare(string(a),string(b)) == 0
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

func transpost(m [][]rune)[][]rune{
  y := len(m)
  x := len(m[0])
  var result [][]rune
  for i:=0; i<x; i++{
    temp := make([]rune, y)
    result = append(result, temp)
  }

  for i:=0; i<len(m); i++{
    for j:=0; j<len(m[i]); j++{
      result[j][i]=m[i][j]
    }
  }
  return result
}

func day13a(m [][]string){
  result := 0
  for i := range m{
    matrix := convert2Rune(m[i])
    result += findReflection(matrix)
  }
  fmt.Println(result)
}

func day13b(m [][]string){
  result := 0
  for i := range m{
    matrix := convert2Rune(m[i])
    result += findSmudgeReflect(matrix)
  }
  fmt.Println(result)
} 

func findSmudgeReflect(m [][]rune)int{
  result := 0
  h, foundH := findSmudgeMirrorPosition(m)
  result = h*100
  if !foundH{
    tm := transpost(m)
    v, foundV := findSmudgeMirrorPosition(tm)
    if !foundV{
      panic("Found neither.")
    }
    result += v
  }
  return result
}

func findSmudgeMirrorPosition(m [][]rune)(int,bool){
  for i:= 0; i<len(m)-1; i++{
    smudged := smudgedEqual(m[i],m[i+1], 0)
    if smudged <= 1{
      //2 pointers to loop through the list
      for j,k:=i-1,i+2; j>=0 && k<len(m) ; j,k = j-1,k+1{ 
        smudged = smudgedEqual(m[j],m[k],smudged)
        // fmt.Println(i,j,k, smudged)
        if smudged >1{
          break
        }
      }
      if smudged == 1{
        return i+1, true
      }
    }
  }
  return 0, false
}

func smudgedEqual(a []rune, b []rune, smudged int)int{
  for i := range a{
    if a[i]!=b[i]{
      smudged++
    }
  }
  return smudged
}
