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
  seedData :=readFileToSeed(filePath)
  day05a(seedData)
}

func readFileToSeed(filePath string) *seedDataContainer{
  readFile, err := os.Open(filePath)
  if err != nil{
    panic(err)
  }
  defer readFile.Close()

  seedData := seedDataContainer{}

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
  //input seeds value
  fileScanner.Scan()
  seedData.seeds = parseSeeds(fileScanner.Text())
  //the empty line
  fileScanner.Scan()
  //scan the rest of the input files
  seedData.seed2soil = parseMapper(fileScanner) 
  seedData.soil2fert = parseMapper(fileScanner) 
  seedData.fert2water = parseMapper(fileScanner) 
  seedData.water2light = parseMapper(fileScanner)
  seedData.light2temp = parseMapper(fileScanner)
  seedData.temp2humid = parseMapper(fileScanner)
  seedData.humid2location = parseMapper(fileScanner)

  return &seedData
}

type seedDataContainer struct{
  seeds []int64
  seed2soil []mapper
  soil2fert []mapper
  fert2water []mapper
  water2light []mapper
  light2temp []mapper
  temp2humid []mapper
  humid2location []mapper
}

type mapper struct{
  source int64
  destination int64
  length int64
}

func parseMapper(fileScanner *bufio.Scanner) []mapper{
  //scan & drop the title
  fileScanner.Scan()

  var result []mapper
  for fileScanner.Scan(){
    text := fileScanner.Text()
    if text == ""{
      break
    }

    //scan the three numbers
    numsString:= strings.Fields(text)
    var numsLine []int64 
    for _, n := range numsString{
      num, err := strconv.ParseInt(n,10,64)
      if err != nil{
        panic(err)
      }
      numsLine = append(numsLine,num)
    }
    m := mapper{
      destination:numsLine[0],
      source: numsLine[1],
      length: numsLine[2],
    }
    result = append(result, m)
  }
  return result
}

func parseSeeds(s string)[]int64{
  var result []int64
  temp := strings.Split(s, ":")
  seedString := strings.Fields(temp[1])
  for _, v := range seedString{
    seed, err := strconv.ParseInt(v,10,64)
    if err != nil{
      panic(err)
    }
    result = append(result, seed)
  }
  return result
}

func day05a(seedData *seedDataContainer){
  var locationList []int64
  for _, s := range seedData.seeds{ 
    r := traceMap(traceMap(traceMap(traceMap(traceMap(traceMap(
      traceMap(s, &seedData.seed2soil),
        &seedData.soil2fert), 
        &seedData.fert2water),
        &seedData.water2light),
        &seedData.light2temp),
        &seedData.temp2humid),
        &seedData.humid2location)
    locationList = append(locationList, r)
  } 
  fmt.Println(searchLowest(&locationList))
}

func traceMap(preNum int64,m *[]mapper)int64{ 
  result := preNum
  for _,v := range (*m){
    if preNum > v.source && preNum < v.source+v.length{
      result = v.destination + (preNum-v.source)
      break
    }
  } 
  return result
}

func searchLowest(locationList *[]int64) int64{
  result := (*locationList)[0] 
  for _, value := range (*locationList){
    if value < result{
      result = value
    }
  }
  return result 
}
