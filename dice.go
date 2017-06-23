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

func initRandomSeed() {
  rand.Seed(time.Now().UTC().UnixNano())
}
