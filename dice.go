package main

import "time"
import "math/rand"

func rollDice(sides int) int {
  return rand.Intn(sides) + 1
}

func rollD6() int {
  return rollDice(6)
}

func rollD20() int {
  return rollDice(20)
}

func randIndex(sliceLen int) int {
  if sliceLen == 0 {
    panic("Cannot call randIndex on an empty slice!")
  }

  return rand.Intn(sliceLen)
}

func initRandomSeed() {
  rand.Seed(time.Now().UTC().UnixNano())
}
