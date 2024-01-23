package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
//  filePath := "./input02.txt"
  filePath := "./input.txt"
  list := readInput(filePath)
  
  r := day09a(list)
  fmt.Printf("The sum of next steps: %d \n",r)
  r02 := day09b(list)
  fmt.Printf("The sum of previous steps: %d \n",r02)
}

func readInput(filePath string)[][]int{
  readFile, err := os.Open(filePath)
  if err != nil{
    panic(err)
  }
  defer readFile.Close()

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  var steps [][]int
  for fileScanner.Scan(){
    var step []int
    line := strings.Fields(fileScanner.Text())
    for _,num := range line{
      n, err := strconv.Atoi(num)
      if err != nil{
        panic(err)
      }
      step = append(step, n)
    }  
    steps = append(steps, step)
  }
  return steps
}

func day09a(list [][]int)int{
  r := 0
  for _, steps := range list{
    r += predictNext(steps)
  }
  return r
}

func predictNext(steps []int)int{
  if checkFinish(steps){
    return steps[0]
  }
  var n []int 
  for i:= 0; i<len(steps)-1; i++{
    n = append(n, steps[i+1]-steps[i])
  }
  r := predictNext(n)
  return r+steps[len(steps)-1]
}

func checkFinish(steps []int)bool{
  for _,v := range steps{
    if v != steps[0]{
      return false
    }
  }
  return true
}

func predictPrev(steps []int) int{
  if checkFinish(steps){
    return steps[0]
  }
  var n []int 
  for i:= 0; i<len(steps)-1; i++{
    n = append(n, steps[i+1]-steps[i])
  }
  r := predictPrev(n)
  return steps[0]-r
}

func day09b(list [][]int)int {
  result := 0
  for _,steps := range list{
    result += predictPrev(steps)
  }
  return result
}
