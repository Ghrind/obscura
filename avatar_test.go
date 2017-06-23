package main

import "testing"
import "fmt"

func expectModFromStat(t *testing.T, stat int, expectedMod int) {
  value := modFromStat(stat)
  if value != expectedMod {
    t.Error(fmt.Sprintf("Expected %d to have a mod of %d, got %d", stat, expectedMod, value))
  }
}

func TestModFromStat(t *testing.T) {
  expectModFromStat(t, 10, 0)
  expectModFromStat(t, 9, -1)
  expectModFromStat(t, 8, -1)
  expectModFromStat(t, 11, 0)
  expectModFromStat(t, 12, 1)
  expectModFromStat(t, 13, 1)
}
