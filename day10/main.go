package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
//  filePath := "./input03.txt"
// filePath := "./input02.txt"
  filePath := "./input.txt"
  maze := readData(filePath) 
  day10a(maze)
}

func readData(filePath string)[]string{
  readFile, err := os.Open(filePath)
  if err != nil{
    panic(err)
  }
  defer readFile.Close()

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  var maze []string
  for fileScanner.Scan(){
    l := fileScanner.Text()
    maze = append(maze, l)
  }
  return maze
}

type point struct{
  y int
  x int
}

func (a *point)add(b point)point{
  c:=point{} 
  c.x = a.x+b.x
  c.y = a.y+b.y
  return c
}

func (a *point)equalPosition(b point)bool{
  return a.x==b.x && a.y==b.y
}

func findStart(maze []string)(point,error){
  var start point 
  for y,l := range maze{
    for x,c := range l{
      if c == SPipe{
        return point{y,x},nil 
      }
    }
  }
  return start, fmt.Errorf("Input error, didn't find starting point.\n")
}

func memoizedStart(maze []string) func()point{
  var start point
  alreadycomputed := false
  return func() point{
    if !alreadycomputed{
      var err error
      start,err = findStart(maze)
      if err != nil{
        panic(err)
      }
      alreadycomputed = true
    }
    return start
  }
} 

func atStart(current point, maze []string)bool{
  startPoint := memoizedStart(maze) 
  return current.equalPosition(startPoint())
}

//scan surrounding to see where can move next 
func findStartNext(start point, maze []string)point{
  n := maze[start.y-1][start.x]
  if (n == NSPipe || n == SWPipe || n == SEPipe){
    return point{y:start.y-1,x:start.x}
  }
  e := maze[start.y][start.x+1]
  if (e == EWPipe || e == NWPipe || e == SWPipe){
    return point{y:start.y, x:start.x+1}
  }
  s := maze[start.y+1][start.x]
  if (s == NSPipe || s == NEPipe || s == NWPipe){
    return point{y:start.y+1, x:start.x}
  }
  w := maze[start.y][start.x-1] 
  if w == EWPipe || w == NEPipe || w == SEPipe{
    return point{y:start.y, x:start.x-1}
  }
  return point{y:start.y+1, x:start.x}
}

var EWPipePosition = []point{{0,-1},{0,1}}
var NSPipePosition = []point{{1,0},{-1,0}}
var NEPipePostion = []point{{-1,0},{0,1}}
var NWPipePostion = []point{{-1,0},{0,-1}}
var SWPipePostion = []point{{1,0},{0,-1}}
var SEPipePostion = []point{{1,0},{0,1}}

type Pipe rune
const (
  NSPipe = '|'
  EWPipe = '-'
  NEPipe = 'L'
  NWPipe = 'J'
  SWPipe = '7'
  SEPipe = 'F'
  Ground = '.'
  SPipe = 'S'
)

func nextPosition(prev, cur point, maze []string)point{ 
  currentPipe := PipeFactory(cur,maze)
  adjustPoint := cur.add(currentPipe[0])
  if prev.equalPosition(adjustPoint){
    adjustPoint = cur.add(currentPipe[1])
  }
  return adjustPoint
}

func PipeFactory(cur point, maze []string)[]point{
  var currentPipe []point
  switch maze[cur.y][cur.x]{
  case NSPipe:
    currentPipe = NSPipePosition 
  case EWPipe:
    currentPipe = EWPipePosition 
  case NEPipe:
    currentPipe = NEPipePostion
  case NWPipe:
    currentPipe = NWPipePostion
  case SWPipe:
    currentPipe = SWPipePostion
  case SEPipe:
    currentPipe = SEPipePostion
  }
  return currentPipe
}

func day10a(maze []string){
  path := calPath(maze)
  fmt.Printf("The furthest Point is %d away.\n",len(path)/2)
}

func calPath(maze []string)[]point{
  startPoint := memoizedStart(maze)
  next := findStartNext(startPoint(), maze)
  var path []point
  path = recurPath(startPoint(), next, maze, path)
  return path
}

func recurPath(prev, current point, maze []string, path []point)[]point{
  path = append(path, current)
  if atStart(current, maze){
    return path
  }
  next := nextPosition(prev, current, maze)
  return recurPath(current, next, maze, path)
}
