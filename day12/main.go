package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//filePath := "./input02.txt"
  filePath := "./input.txt"
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
	if !checkValidRecord(record, config) {
		return 0
	}
	result := 0
	//strip the leading working char
	if record[0] == working {
		if len(config) == 0 {
			return 1
		}
		return handleRecord(record[1:], config)
	}
	if record[0] == damaged {
		if validContinuous(record, config) {
			if len(config) == 1 {
				return 1
			}
			return handleRecord(record[config[0]+1:], config[1:])
		}
		return 0
	}
	if record[0] == unknown {
		temp := 0
		//asssume working
		// fmt.Printf(".%s %v %d\n", record[1:], config, temp)
		temp += handleRecord(record[1:], config)
		// fmt.Println(temp)
		//assume broken
		if validContinuous(record, config) {
			if len(config) == 1 {
          return temp+1
			}
			// fmt.Printf("#%s %v %d\n", record[config[0]+1:], config[1:], temp)
			temp += handleRecord(record[config[0]+1:], config[1:])
			// fmt.Println(temp)
		}
		result += temp
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
  result := 0
  for _, i := range hs{
    r := handleRecord(i.record, i.config)
   // fmt.Println(r)
    result += r
  }
  fmt.Println(result)
}
