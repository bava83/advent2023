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
  fmt.Printf("Day04b total cards:%d\n",day04b(listCard))
}

func day04a(list []Card)int{
  sum := 0
  wg := sync.WaitGroup{}
  for x:=0; x<len(list); x++{
    wg.Add(1)
    go list[x].findWinning(&wg)
    wg.Wait()
    list[x].calPoints()
    sum += list[x].winPoint 
  }
  return sum
}

func day04b(list []Card)int{
  result := 0
  setCopies(list)
  for _,c := range list{
    result += c.copies
  }
  return result
}



type Card struct{
  winNum []int
  ownNum []int
  matchCount int
  winPoint int
  copies int 
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

    c := Card{copies: 1}
    // var c Card
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
  c.matchCount = 0
  ch := make (chan bool)
  for _, wn := range c.winNum{
    go containWinnings(c.ownNum, wn, ch)
    if <- ch{
      c.matchCount++
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
  if c.matchCount == 0{
    c.winPoint = 0
  }
  c.winPoint = int(math.Pow(2, float64(c.matchCount-1)))
}

func setCopies(list []Card){
  for i, c := range list{
    for x:=0; x<c.matchCount; x++{
      for y:=0; y<c.copies; y++{
        list[i+x+1].copies++
      }
    }    
  }
}
