package main

import (
	"fmt"
)


type data struct{
  time int
  dist int
}

func main(){
  /* var input = []data{
    {
      time : 7,
      dist : 9,
  },{
    time : 15,
    dist : 40,
  },{
      time : 30,
      dist : 200,
    },
  } */
  var input = []data{
    {
      time : 57,
      dist: 291,
    },{
      time: 72,
      dist: 1172,
    },{
      time: 69,
      dist: 1176,
    },{
      time: 92,
      dist: 2026,
    },
  }
  var input02 = data{
    time: 57726992,
    dist: 291117211762026,
  }
  day06a(input)
  day06b(input02)
}

func day06a(input []data){
  result := 1
//max-min+1
  for _,v := range input{
    maxHold := maxHolding(v.time, v.dist)
    minHold := minHolding(v.time, v.dist) 
    result *= maxHold-minHold+1
  }
  fmt.Println(result)
}

func day06b(input data){
  maxHold := maxHolding(input.time, input.dist)
  minHold := minHolding(input.time, input.dist)
  result := maxHold-minHold+1 
  fmt.Println(result)
}

func minHolding(time int, distRecord int)int{
  i:=0
  for ; i<time; i++{
    speed := i
    distTravel := (time-i)*speed
    if distTravel > distRecord {
      break
    }  
  }
  return i
}

func maxHolding(time int, distRecord int)int{
  i:=time
  for ; i>=0; i--{
    speed := i
    distTravel := (time-i)*speed
    if distTravel > distRecord{
      break
    }
  }
  return i
}
