package main

import "fmt"
import "os"
import "math/rand"
import "time"

func main() {
  fmt.Printf("Welcome to Crawler!\n")

  rand.Seed( time.Now().UTC().UnixNano())

  player1, err := load()

  if err != nil {
    fmt.Println(err)
  }

  if player1.Name == "" {
    player1 := avatar{}
    player1.Name = "unknown"
    player1.Class = "peon"
    player1 = rollAvatar(player1)
  }

  fmt.Printf("Here's your avatar:\n")
  showAvatar(player1)

  player1 = showPlayerMenu(player1)

  save(player1)

  os.Exit(0)

}
