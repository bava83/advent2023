package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
//  filePath := "./input.txt"
  filePath := "./input02.txt"
  hs := readFile(filePath)
  day12(hs)
}

func readFile(filePath string)[]hotSpring{
  readFile, err := os.Open(filePath)
  if err!= nil{
    panic(err)
  }
  defer readFile.Close()

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  var list []hotSpring
  for fileScanner.Scan(){
    var hs hotSpring
    l := strings.Split(fileScanner.Text(), " ")
    hs.record = l[0]
    c := strings.Split(l[1], ",")
    for _, i := range c{
      num, err := strconv.Atoi(i)
      if err != nil{
        panic(err)
      }
      hs.config = append(hs.config, num)
    }
    list = append(list, hs)
  }
  return list
}

func day12(hs []hotSpring){
  //testing
  
  r := []int{1,1,3}
  s01 := handleRecords("???.###", r)
  fmt.Println(s01)


}

type hotSpring struct{
  record string
  config []int
}

const (
  damaged = '#'
  working = '.'
  unknown = '?'
)


func handleRecords(record string, configs []int)int{
  result := 0
  //strip the prefix that's working
  if record[0] == working{
    return handleRecords(record[1:], configs)
  }
  var re []int
  start := 0
  for i:=0; i<len(configs)-1; i++{
    temp := singlePattern(record[start:(start+configs[i])],configs[i])
    start += configs[i]+1
    re = append(re, temp)
  }
  temp := singlePattern(record[start:(start+configs[len(configs)-1])], configs[len(configs)-1])
  re = append(re, temp)
  fmt.Println(re)
  return result
}

func singlePattern(record string, config int)int{
  r := 0
  if len(record)==0{
    if config == 0{
      return 1
    }
    return 0
  }
  if len(record)==1{
    return handleBase(record,config)
  }
  //for case len > 1
  switch record[0]{
    case working:
      return singlePattern(record[1:],config)
    case damaged:
      //wrong
      return handleContinuous(record, config)
    case unknown:
      u := 0
      //assume first char is working
      u += singlePattern(record[1:], config)
      //assume first char is damage 
      u += handleContinuous(record, config)
      r = u
  }
  return r
}

//handle when length is 1
func handleBase(record string, config int)int{
  switch record[0]{
    case working:
      if config == 0{
        return 1
      }
      return 0
    case damaged:
      if config == 1{
        return 1
      } 
      return 0
    case unknown:
      if config == 1 || config == 0{
        return 1
      }
      return 0
  }
  return 0
}

func handleContinuous(record string, config int)int{
  if len(record) == 0 || config > len(record){
    return 0
  }

  for i:=0; i<len(record) && i<config; i++{
    if record[i] == working{
      return 0
    } 
  }
  //check what's after the damaged, cannot be damaged
  if len(record)> config{
    if record[config] == damaged{
      return 0
    }
  }
  return 1
}
