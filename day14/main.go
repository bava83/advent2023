package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
  filePath := "./input.txt"
//  filePath := "./input02.txt"
//  filePath := "./input03.txt"
  m := readFilesToMatrix(filePath)
//  day14a(m)
  day14b(m)
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

type platform [][]rune

const (
  roundRock = 'O'
  cubeRock = '#'
  emptySpace = '.'
)

func day14a(m [][]rune){
  result := 0
  p := platform(m)
  p.tiltNorth()
  result = p.calLoad()
  fmt.Println(result)
}

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

func day14b(m [][]rune){
  var p platform
  p = platform(m)
  //p = p.rotateR90()
  //p = p.cycle()
  r := loopCycle(p)
  fmt.Printf("After 1 billion cycle, the load is %d\n",r)
}

func (p platform) print(){
  for _,i := range p{
    fmt.Println(string(i))
  }
  fmt.Println("")
}

func (p platform) cycle()platform{
  p.tiltNorth()
  //west
  p = p.rotateR90()
  p.tiltNorth()
  //south
  p = p.rotateR90()
  p.tiltNorth()
  //east
  p = p.rotateR90()
  p.tiltNorth()
  //rotate back to north
  return p.rotateR90()
}

func (p platform)rotateR90()platform{
  result := make(platform, len(p[0]))
  for i:=0; i <len(p[0]); i++{
    result[i] = make([]rune, len(p))
  }

  for i:=0; i<len(p); i++{
    for j:=0; j<len(p[i]);j++{
      result[j][len(p)-1-i] = p[i][j]
    }  
  }
  return result
}


type platformData struct{
  cycle int
  load int
}

func loopCycle(p platform)int{
  repeat := 1000000000
  hs := make(map[string]platformData)
  result := 0

  modSize := 0
  var s string
  for i:=1; i<repeat; i++{
    p = p.cycle()
    s = platform2String(p)
    _, ok := hs[s]
    //found the repeat pattern
    if ok{
      modSize = i-hs[s].cycle
      break
    }
    tmp := platformData{i,p.calLoad()}
    hs[s] = tmp 
  }

  //the mod cycle happens after some initial cycle
  //so 1bil repeats need to minus it first, then add back after 
  modValue := (repeat - hs[s].cycle)%modSize + hs[s].cycle
  for _,v := range hs{
    if v.cycle == modValue{
      result = v.load
    } 
  }
  return result
}

func platform2String(p platform)string{
  var s string
  for i:=0; i<len(p); i++{
    s += string(p[i])
  }
  return s
}
