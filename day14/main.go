package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
  filePath := "./input.txt"
//  filePath := "./input02.txt"
  m := readFilesToMatrix(filePath)
  day14a(m)
}

func readFilesToMatrix(filePath string)[][]rune{
  readFile, err := os.Open(filePath)
  if err != nil{
    panic(err)
  }
  defer readFile.Close()
  
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  var matrix [][]rune
  for fileScanner.Scan(){
    l := fileScanner.Text()
    var line []rune
    for _, c := range l{
      line = append(line, c)
    }
    matrix = append(matrix, line)
  }

  return matrix
}

func day14a(m [][]rune){
  result := 0
  p := platform(m)
  p.tiltNorth()
  result = p.calLoad()
  fmt.Println(result)
}

type platform [][]rune

const (
  roundRock = 'O'
  cubeRock = '#'
  emptySpace = '.'
)

func (p platform) tiltNorth(){
  for i:=0; i<len(p); i++{
    for j:=0; j<len(p[i]); j++{
      if p[i][j] == roundRock{
       k := i-1 
        for ; k>=0; k--{
          if p[k][j]!=emptySpace{
            break
          }
        } 
        if i!=k+1{
          p[i][j] = emptySpace
          p[k+1][j] = roundRock
        }
      }
    }
  }
}

func (p platform) calLoad()int{
  result := 0
  for i,load := 0, len(p); i<len(p); i,load = i+1, load-1 {
    for j:=0; j<len(p[i]); j++{
      if p[i][j] == roundRock{
        result += load
      } 
    }  
  }
  return result 
}
