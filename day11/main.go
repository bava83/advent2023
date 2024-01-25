package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){
  filePath := "./input.txt"
//  filePath := "./input02.txt"
  u := readFile(filePath)
// remember u is a pointer to the universe, so day11a will mutate u 
//  day11a(u)
  day11b(u)
}

func readFile(filePath string)universe{
  readFile, err := os.Open(filePath)
  if err != nil{
    panic(err)
  }
  defer readFile.Close() 

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
  
  var galaxy []point
  row := 0
  var l string
  for fileScanner.Scan(){
    l = fileScanner.Text()
    for col:=0; col<len(l); col++{
      var p point
      if l[col] == '#'{
        p.x = col
        p.y = row
        galaxy = append(galaxy, p)
      }
    } 
    row++
  }

  u := universe{galaxy,row, len(l)}
  return u 
}

type point struct{
  y int
  x int
}

type universe struct{
  galaxy []point
  row int
  column int 
}

func day11a(u universe){
  u.expendUniverse(2)
  r := u.calTotalDistance()
  fmt.Println(r)
}

func day11b(u universe){
  u.expendUniverse(1000000)
  r := u.calTotalDistance()
  fmt.Println(r)
}


func (u *universe)expendUniverse(emptySize int){  
  emptyRows := findEmptyRows(u)
  emptyCols := findEmptyCols(u)
  u.expRows(emptyRows, emptySize)
  u.expCols(emptyCols, emptySize)
}

func findEmptyRows(u *universe)[]int{
  //find the empty rows
  var rows4Exp []int
  for row:=0; row<u.row; row++{
    needed := true
    for i:=0; i<len(u.galaxy); i++{
      //rows not empty look for next row
      if u.galaxy[i].y == row{
        needed = false 
        break
      }
    }
    if needed{
      rows4Exp = append(rows4Exp,row)
    }
  }
  return rows4Exp
}

func findEmptyCols(u *universe)[]int{
  var col4Exp []int
  for col:=0; col<u.column; col++{
    needed := true
    for i:=0; i<len(u.galaxy); i++{
      if u.galaxy[i].x == col{
        needed = false
        break
      }
    }
    if needed{
      col4Exp = append(col4Exp, col)
    }
  } 
  return col4Exp
}

func (u *universe) expRows(emptyRows []int, emptySize int){
  for ri, row := range emptyRows{
    for i:=0; i<len(u.galaxy); i++{
      if u.galaxy[i].y > row{
        u.galaxy[i].y += emptySize-1
      }
    }
    //since we moved the galaxy, we also need to move the empty lines
    //eg. galaxy row 7 is increase to 9, which is located after emptyrow 8
    //so we don't want the original 7 count as after 8, so we have to increase the empty row as well
    for rj:=ri+1; rj<len(emptyRows); rj++{
      emptyRows[rj]+=emptySize-1
    }
  }
}

func (u *universe) expCols(emptyCols []int, emptySize int){
  for ci, col := range emptyCols{
    for i:=0; i<len(u.galaxy); i++{
      if u.galaxy[i].x > col{
        u.galaxy[i].x += emptySize-1      
      }
    }
    for cj:=ci+1; cj<len(emptyCols);cj++{
      emptyCols[cj]+=emptySize-1
    } 
  }
}

func calDistance(a point, b point)int{
  result := 0
  result += int(math.Abs(float64(b.y)-float64(a.y)))
  result += int(math.Abs(float64(b.x)-float64(a.x)))
  return result
}

func (u *universe)calTotalDistance()int{
  result := 0
  for i:=0; i<len(u.galaxy);i++{
    for j:=i+1; j<len(u.galaxy); j++{
      result += calDistance(u.galaxy[i],u.galaxy[j])
      //fmt.Printf("i:%v j:%v r:%d \n",u.galaxy[i], u.galaxy[j],r)
    }
  }
  return result
}
