package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "./input02.txt"
  //filePath := "./input.txt"
	hs := readFile(filePath)
	day12(hs)
}

func readFile(filePath string) []hotSpring {
	readFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var list []hotSpring
	for fileScanner.Scan() {
		var hs hotSpring
		l := strings.Split(fileScanner.Text(), " ")
		hs.record = l[0]
		c := strings.Split(l[1], ",")
		for _, i := range c {
			num, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			hs.config = append(hs.config, num)
		}
		list = append(list, hs)
	}
	return list
}

func day12(hs []hotSpring) {
	//testing
  day12a(hs)
}

type hotSpring struct {
	record string
	config []int
}

const (
	damaged = '#'
	working = '.'
	unknown = '?'
)

func validContinuous(record string, config []int) bool {
	if len(record) == 0 || config[0] > len(record) {
		return false
	}

	for i := 0; i < len(record) && i < config[0]; i++ {
		if record[i] == working {
			return false
		}
	}
	//check what's after the damaged, cannot be damaged
	if len(record) > config[0] {
		if record[config[0]] == damaged {
			return false
		}
	}
	return true
}

func handleRecord(record string, config []int) int {
  if record == ""{
    if len(config) == 0{
      return 1
    }
    return 0
  }
  if len(config) == 0{
    if strings.Contains(record, string(damaged)){
      return 0
    }
    return 1
  }

	result := 0

  if record[0] == working || record[0] == unknown{
    result += handleRecord(record[1:], config)
  }

  if record[0] == damaged || record[0] == unknown{
    if config[0] <= len(record) && !(strings.Contains(record[:config[0]],string(working))) && (config[0] == len(record) || record[config[0]] != damaged){
      result += handleRecord(record[config[0]+1:], config[1:])
    }
  }
	return result
}

func checkValidRecord(record string, config []int) bool {
	allConfigSum := getConfigSum(config)
	if allConfigSum > len(record) {
		return false
	}
	return true
}

func getConfigSum(config []int) int {
	result := 0
	for _, c := range config {
		result += c
	}
	return result
}

func day12a(hs []hotSpring){
/*   result := 0
  for _, i := range hs{
    r := handleRecord(i.record, i.config)
   // fmt.Println(r)
    result += r
  }
  fmt.Println(result) */
  s01 := "???.###"
  a01 := []int{1,1,3}
  t := handleRecord(s01,a01)
  fmt.Println(t)
}
