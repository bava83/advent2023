package main

import "testing"

func TestGetFirstDigit(t *testing.T){
  var tests = []struct{
    input string
    expected int
  }{
    {"1abc2", 1},
    {"pqr3stu8vwx",3},
    {"a1b2c3d4e5f",1},
    {"treb7uchet",7},
  }

  for _, tt := range tests{
    actual, _ := getFirstDigit(tt.input)
    if actual != tt.expected{
      t.Errorf("getFirstDigit(%s) = %d; expected: %d", tt.input, actual, tt.expected)
    }
  }
}


func TestGetLastDigit(t *testing.T){
  var tests = []struct{
    input string
    expected int
  }{
    {"1abc2", 2},
    {"pqr3stu8vwx",8},
    {"a1b2c3d4e5f",5},
    {"treb7uchet",7},
  }

  for _, tt := range tests{
    actual, _ := getLastDigit(tt.input)
    if actual != tt.expected{
      t.Errorf("getLastDigit(%s) = %d; expected: %d", tt.input, actual, tt.expected)
    }
  }
}
