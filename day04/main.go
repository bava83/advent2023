package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main(){
//  filePath := "./input02.txt"
  filePath := "./input.txt"
  listCard := readFileToCard(filePath)
  fmt.Printf("Day04a total points:%d\n",day04a(listCard))
}

func day04a(list []Card)int{
  sum := 0
  wg := sync.WaitGroup{}
  for _, c := range list{
    wg.Add(1)
    go c.findWinning(&wg)
    wg.Wait()
    c.calPoints()
    sum += c.winPoint 
  }

  return sum
}

type Card struct{
  winNum []int
  ownNum []int
  winCount int
  winPoint int
}

func readFileToCard(filePath string) []Card{
  readFile, err := os.Open(filePath)
  if err !=  nil{
    panic(err)
  }
  defer readFile.Close()

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  var cardList []Card
  for fileScanner.Scan(){
    l := fileScanner.Text()
    line := strings.Split(l, ":")
    entryString := strings.Split(line[1], "|")
    winString := strings.Fields(entryString[0])
    numString := strings.Fields(entryString[1])

    var c Card
    c.winNum = entry2Int(winString)
    c.ownNum = entry2Int(numString)
    cardList = append(cardList, c)
  }
  return cardList
}

func entry2Int(str []string)[]int{
  var numbers []int
  for _,s := range str{
    n, err := strconv.Atoi(s)
    if err != nil{
      panic(err)
    }
    numbers = append(numbers, n)
  }
  return numbers
}

//find how many winning numbers
func (c *Card)findWinning(wg *sync.WaitGroup) {
  c.winCount = 0
  ch := make (chan bool)
  for _, wn := range c.winNum{
    go containWinnings(c.ownNum, wn, ch)
    if <- ch{
      c.winCount++
    } 
  }
  defer wg.Done()
}

func containWinnings(list []int, target int, ch chan bool){
  result := false 
  for _, v := range list{
    if target == v{
      result = true
      break
    }
  } 
  ch <- result
}

func (c *Card) calPoints(){
  if c.winCount == 0{
    c.winPoint = 0
  }
  c.winPoint = int(math.Pow(2, float64(c.winCount-1)))
}
