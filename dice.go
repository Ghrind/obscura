package main

import "time"
import "math/rand"

func RandInt(min int, max int) int {
  return rand.Intn(max - min + 1) + min
}

func randIndex(sliceLen int) int {
  if sliceLen == 0 {
    panic("Cannot call randIndex on an empty slice!")
  }

  return rand.Intn(sliceLen)
}

func InitRandomSeed() {
  rand.Seed(time.Now().UTC().UnixNano())
}
