package main

import "fmt"
import "os"
import "bufio"
import "math/rand"
import "strconv"
import "encoding/json"
import "io/ioutil"
import "time"

type avatar struct {
  Name string
  Class string
  Str int
  Dex int
  Con int
  Int int
  Wis int
  Cha int
}

func rollAvatar(avatar avatar) avatar {
  avatar.Str = rand.Intn(6) + rand.Intn(6) + 6 + 2
  avatar.Dex = rand.Intn(6) + rand.Intn(6) + 6 + 2
  avatar.Con = rand.Intn(6) + rand.Intn(6) + 6 + 2
  avatar.Int = rand.Intn(6) + rand.Intn(6) + 6 + 2
  avatar.Wis = rand.Intn(6) + rand.Intn(6) + 6 + 2
  avatar.Cha = rand.Intn(6) + rand.Intn(6) + 6 + 2

  return avatar
}

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

func showPlayerMenu(avatar avatar) avatar {
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
      classes := []string{"warrior", "hunter"}
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

func save(avatar avatar) {
  content, _ := json.Marshal(avatar)
  err := ioutil.WriteFile("/tmp/crawler.json", content, 0644)
  if err != nil {
    fmt.Println(err)
  }
}

func load() (avatar, error) {
  avatar := avatar{}
  content, err := ioutil.ReadFile("/tmp/crawler.json")
  if err != nil {
    return avatar, err
  }
  err = json.Unmarshal(content, &avatar)
  if err != nil {
    return avatar, err
  }
  return avatar, nil
}

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
