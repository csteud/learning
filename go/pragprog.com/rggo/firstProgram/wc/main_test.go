package main

import (
  "bytes"
  "testing"
)

func TestCountWords(t *testing.T) {
  b := bytes.NewBufferString("word1 word2 word3 word4\n")

  exp := 4 
  
  res := count(b, false, false)

  if res != exp {
    t.Errorf("Expected %d, got %d in-stead.\n", exp, res)
  }
}

func TestCountLines(t *testing.T) {
  b := bytes.NewBufferString("line1 word\n line2 word\n")

  exp := 2 

  res := count(b, true, false)

  if res != exp {
    t.Errorf("Expected %d, got %d in-stead.\n", exp, res)
  }
}

func TestCountBytes(t *testing.T) {
  b := bytes.NewBufferString("first second")

  exp := 12 

  res := count(b, false, true)

  if res != exp {
    t.Errorf("Expected %d, got %d in-stead.\n", exp, res)
  }
}
