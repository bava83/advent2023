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
  r02 := day08b(command,m)
  fmt.Printf("It takes %d steps to reach __Z.\n",r02)
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
  return findSteps(command, m, "AAA", checkMatch)
}

type checkerFunction func(string) bool
func checkMatch(s string)bool{
  return s=="ZZZ" 
}

func findSteps(command string, m map[string]node, start string, checker checkerFunction) int{
  instruction := start
  steps := 0
  for {
    for _, c := range command{
      if checker(instruction){
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

type nodeSteps struct{
  node string
  steps int
}

func day08b(command string, m map[string]node)int{
  checkList := searchStartingNodes(m)
  commonSteps := 0
  var listSteps []int
  for i := range checkList{
    checkList[i].steps = findSteps(command, m, checkList[i].node, checkEnd)
    listSteps = append(listSteps, checkList[i].steps)
  }
  commonSteps = LCM(listSteps[0], listSteps[1], listSteps[2:]...)
  
  return commonSteps
}

//check if string match with __Z
func checkEnd(s string)bool{
  return s[len(s)-1] == 'Z'
}

/* func checkEndWithZ(l []string) bool{
  for _, s := range l{
    if s[len(s)-1] != 'Z'{
      return false
    }
  }
  return true
} */

//find the list that starts with __A 
func searchStartingNodes(m map[string]node)[]nodeSteps{
  var list []nodeSteps
  for k := range m{
    if k[len(k)-1] == 'A'{
      c := nodeSteps{node: k, steps: 0}
      list = append(list, c)
    }
  }
  return list
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}

