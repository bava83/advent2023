package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
//  filePath := "./input02.txt"
  filePath := "./input.txt"
  command, m := readFile(filePath)
  r := day08a(command, m)
  fmt.Printf("It takes %d steps to reach ZZZ.\n",r)
}

func readFile(filePath string) (string, map[string]node){
  readFile, err := os.Open(filePath)
  if err != nil{
    panic(err)
  }
  defer readFile.Close()

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
  //scan the command  
  fileScanner.Scan()
  command := fileScanner.Text()
  //scan the empty line
  fileScanner.Scan()
  //scan the input data
  m := make(map[string]node)
  for fileScanner.Scan(){
    l := fileScanner.Text()
    line := strings.Split(l, " = ") 
    n := node{parent: line[0]}
    childrenText := strings.Split(strings.Trim(line[1], "()"), ", ")
    n.left = childrenText[0]
    n.right = childrenText[1]
    m[line[0]] = n
  }
  return command, m
}

type node struct{
  parent string
  left string
  right string
}


func day08a(command string, m map[string]node) int{
  instruction := "AAA"
  steps := 0
  for {
    for _, c := range command{
      if instruction == "ZZZ"{
        return steps 
      }
      switch c{
      case 'L':
        instruction = m[instruction].left
      case 'R':
        instruction = m[instruction].right 
      }
      steps++
    } 
  }
}
