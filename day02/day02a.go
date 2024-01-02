package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
	"strconv"
	"strings"
)

func day02a(){
  filePath := "./input.txt"
  readFile, err := os.Open(filePath)
  if err != nil {
    panic(err)
  }
  defer readFile.Close()
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  var games []gameSet
  for fileScanner.Scan(){
    s := fileScanner.Text()
    g := gameSet{valid: false}
    g.parseGameInput(s)
    games = append(games, g)
  }

}

type gameSet struct {
  id int
  drawList []draw
  valid bool
}

type draw struct{
  red int
  green int
  blue int
}

func (g *gameSet) parseGameInput(s string) {
  inputArray := strings.Split(s, ":")
  var err error
  //input game id
  g.id, err = strconv.Atoi(strings.Split(inputArray[0], " ")[1]) 
  if err != nil{
    panic(err)
  }
 
  //split the 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green into draw
  gameInputString := strings.Split(inputArray[1], ";") 
  for _, drawString := range gameInputString{
    d := draw{}
    //split the 3 blue, 4 red 
    cubeString := strings.Split(drawString, "," )
    for _, cube := range cubeString{
      //split the 3 blue
      cubeInput := strings.Fields(cube)
      n, err := strconv.Atoi(cubeInput[0])
      if err != nil {
        panic(err)
      }
      //parse the 3, blue
      d.parseCubeData(n, cubeInput[1])
    }
    g.drawList = append(g.drawList, d)
  }
}

func (d *draw) parseCubeData(n int, color string){
  switch color{
  case "red":
    d.red = n
  case "green":
    d.green = n
  case "blue":
    d.blue = n 
  }
}
