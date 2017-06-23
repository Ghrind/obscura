//
// ui.go
//
// Contains all the UI related functions
//
// Naming:
//
// * 'show', not 'display'
//
package main

import "bufio"
import "strconv"
import "fmt"
import "os"

func showAvatar(avatar avatar) {
  fmt.Printf("%s (%s)\n", avatar.Name, avatar.Class)
  fmt.Printf("STR: %d\n", avatar.Str)
  fmt.Printf("DEX: %d\n", avatar.Dex)
  fmt.Printf("CON: %d\n", avatar.Con)
  fmt.Printf("INT: %d\n", avatar.Int)
  fmt.Printf("WIS: %d\n", avatar.Wis)
  fmt.Printf("CHA: %d\n", avatar.Cha)
}

func showMenu() {
  fmt.Printf("(r)eroll, (n)ame the character, change the (c)lass of your character, (q)uit\n")
}

func showClassesMenu(classes []string) {
  for i:=0; i < len(classes); i++ {
    fmt.Printf("%d: %s\n", i, classes[i])
  }
}

func showPlayerMenu(avatar avatar, classes []string) avatar {
  showMenu()

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {

    switch scanner.Text() {
    case "r":
      // Reroll
      avatar = rollAvatar(avatar)
      fmt.Printf("Here's your avatar:\n")
      showAvatar(avatar)
      break
    case "n":
      // Name the character
      fmt.Printf("Give a name to this character (current is '%s'):\n", avatar.Name)
      scanner.Scan()
      avatar.Name = scanner.Text()
      showAvatar(avatar)
      break
    case "c":
      // Change the class
      showClassesMenu(classes)
      scanner.Scan()
      choice, err := strconv.Atoi(scanner.Text())
      if err == nil && choice >= 0 && choice < len(classes) {
        avatar.Class = classes[choice]
      } else {
        fmt.Printf("Bad choice...\n")
      }
      showAvatar(avatar)
    case "q":
      // Quit
      return avatar
      break
    }

    showMenu()

  }

  return avatar
}

