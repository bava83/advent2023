package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main(){
//  filePath := "./input02.txt"
  filePath := "./input.txt"
  cards := readFile(filePath)
  day07a(cards)
}

func readFile(filePath string) []cardBit{
  readFile, err := os.Open(filePath)
  if err != nil{
    panic(err)
  }
  defer readFile.Close()

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  var cardList []cardBit 
  for fileScanner.Scan(){
    text := strings.Fields(fileScanner.Text())
    var c cardBit
    c.hand = input2Card(text[0]) 
    c.bit = input2Bit(text[1]) 
    cardList = append(cardList, c)
  }
  return cardList
}

type cardBit struct{
  hand []int 
  bit int
  t handType
}
type cardListByHand []cardBit 

const (
  T = 10
  J = 11
  Q = 12
  K = 13
  A = 14
)

type handType int
const (
  HIGHC handType = iota
  ONEPAIR
  TWOPAIR
  THREEKIND
  FULLHOUSE
  FOURKIND
  FIVEKIND  
)

func input2Card(s string)[]int{
  var result []int
  for _,c := range s{
    if unicode.IsDigit(c){
      n, err := strconv.Atoi(string(c))
      if err != nil{
        panic(err)
      }
      result = append(result, n)
    } else{
      n, err := convertFaceCard(c)
      if err != nil{
        panic(err)
      }
      result = append(result, n)
    }
  }
  return result
}

func convertFaceCard(c rune)(int,error){
  switch c{
  case 'T':
    return T,nil
  case 'J':
    return J,nil
  case 'Q':
    return Q,nil
  case 'K':
    return K,nil
  case 'A':
    return A,nil
  }
  return 0, fmt.Errorf("Unexpected Char: %c", c)
}

func input2Bit(s string)(int){
  n, err := strconv.Atoi(s)
  if err != nil{
    panic(err)
  }
  return n
}

func initHandsType(cards []cardBit){
  for i:= 0; i<len(cards); i++{
    (cards)[i].setHandType()
  } 
} 

func (cards *cardBit) setHandType(){
  h := make(map[int]int)
  for _, card := range cards.hand{
    if _,ok := h[card]; ok{
      h[card]++
    }else{
      h[card]=1
    }
  }
  threeCard := false  
  twoCard := false
  //default to high card
  cards.t = HIGHC
  for _,v := range h{
    switch v{
    case 5:
      cards.t = FIVEKIND
    case 4:
      cards.t = FOURKIND
    case 3:
      threeCard = true
    case 2:
      //found 1 pair
      if !twoCard{
        twoCard = true
      }else{
        //found 2 pair,set TWOPAIR, but turn off found flag 
        cards.t = TWOPAIR
        twoCard = false
      }  
    }
  }
  //check the flags
  if twoCard && threeCard{
    cards.t = FULLHOUSE
  } else if twoCard && !threeCard{
    cards.t = ONEPAIR
  } else if threeCard && !twoCard{
    cards.t = THREEKIND
  }
}

func (hand01 *cardBit) strongerHand (hand02 *cardBit)bool{
  if hand01.t > hand02.t{
    return true
  } 
  if hand01.t < hand02.t{
    return false
  }
  return hand01.strongerTie(hand02)
}

func (hand01 *cardBit)strongerTie(hand02 *cardBit)bool{
  for i:=0; i<len(hand01.hand);i++{
    if hand01.hand[i] > hand02.hand[i]{
      return true
    }
    if hand01.hand[i] < hand02.hand[i]{
      return false
    } 
  }
  return false 
}

//implement interface for sorting
func (a cardListByHand) Len() int {return len(a)}
func (a cardListByHand) Swap(i,j int) {a[i], a[j] = a[j], a[i]}
func (a cardListByHand) Less(i,j int) bool{
  return !a[i].strongerHand(&a[j])
}

func day07a(cards cardListByHand){
  initHandsType(cards)
  sort.Sort(cardListByHand(cards))
  result := calPoints(cards)
  fmt.Printf("The Winning points: %d \n",result)
}

func calPoints(cards cardListByHand)int{
  points := 0
  for i,card := range cards{
    points += card.bit * (i+1) 
  }
  return points
}

