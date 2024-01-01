package main

import "testing"

func TestGetSpellFirstDigit(t *testing.T){
  var tests = []struct{
    input string
    expected int
  }{
    {"two1nine", 2},
    {"eightwothree",8},
    {"4nineeightseven2",4},
    {"7pqrstsixteen",7},
  }

  for _, tt := range tests{
    actual := getSpellFirstDigit(tt.input)
    if actual != tt.expected{
      t.Errorf("getSpellFirstDigit(%s)=%d, expected %d",tt.input, actual,tt.expected)
    }
  }
}


func TestGetSpellLastDigit(t *testing.T){
  var tests = []struct{
    input string
    expected int
  }{
    {"two1nine", 9},
    {"eightwothree",3},
    {"4nineeightseven2",2},
    {"7pqrstsixteen",6},
  }

  for _, tt := range tests{
    actual := getSpellLastDigit(tt.input)
    if actual != tt.expected{
      t.Errorf("getSpellLastDigit(%s)=%d, expected %d",tt.input, actual,tt.expected)
    }
  }
}
