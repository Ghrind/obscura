//
// ui.go
//
// Contains all the UI related functions
//
// Naming:
//
// * 'show', not 'display'
// * Use 'ask' when getting an input from the user
//
package main

import "bufio"
import "strconv"
import "fmt"
import "os"
import "strings"

func showAvatar(avatar avatar) {
  fmt.Printf("%s (%s)\n", avatar.Name, avatar.Class)
  fmt.Printf("STR: %d\n", avatar.Str)
  fmt.Printf("DEX: %d\n", avatar.Dex)
  fmt.Printf("CON: %d\n", avatar.Con)
  fmt.Printf("INT: %d\n", avatar.Int)
  fmt.Printf("WIS: %d\n", avatar.Wis)
  fmt.Printf("CHA: %d\n", avatar.Cha)
}

func askAction(title string, availableActions []string) string {
  fmt.Println(title)

  scanner := bufio.NewScanner(os.Stdin)
  prompt := fmt.Sprintf("Your choice (%s): ", strings.Join(availableActions, "/"))
  fmt.Printf(prompt)
  for scanner.Scan() {
    input := scanner.Text()
    for i := 0; i < len(availableActions); i++ {
      if availableActions[i] == input {
        return input
      }
    }
    fmt.Printf(prompt)
  }

  return ""
}

func askString(title string, prompt string) string {
  fmt.Println(title)
  fmt.Printf(prompt)

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  return scanner.Text()
}

func askFromList(title string, list []string) string {
  fmt.Println(title)

  for i:=0; i < len(list); i++ {
    fmt.Printf("%d: %s\n", i, list[i])
  }

  prompt := fmt.Sprintf("Your choice (0-%d): ", len(list) - 1)

  fmt.Printf(prompt)

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    choice, err := strconv.Atoi(scanner.Text())
    if err == nil && choice >= 0 && choice < len(list) {
      return list[choice]
    }
    fmt.Printf(prompt)
  }

  return ""
}

func showPlayerMenu(avatar avatar, classes []string) avatar {
  prompt := "(r)eroll, (n)ame the character, change the (c)lass of your character, (q)uit"

  for {
    input := askAction(prompt, []string{"r", "n", "c", "q"})

    switch input {
    case "r":
      // Reroll
      avatar = rollAvatar(avatar)
      fmt.Printf("Here's your avatar:\n")
      showAvatar(avatar)
      break
    case "n":
      // Name the character
      avatar.Name = askString(fmt.Sprintf("Give a name to this character (current is '%s'):", avatar.Name), "New name: ")
      showAvatar(avatar)
      break
    case "c":
      // Change the class
      avatar.Class = askFromList(fmt.Sprintf("Choose a class for your character (current is '%s'):", avatar.Class), classes)
      showAvatar(avatar)
    case "q":
      // Quit
      return avatar
      break
    }

  }

  return avatar
}

func showMeleeMenu(playerAvatar combatAvatar, ennemyAvatar combatAvatar) {

  fmt.Printf("Melee: %s vs %s\n", playerAvatar.name, ennemyAvatar.name)

  showMeleeStatus(playerAvatar, ennemyAvatar)

  prompt := "(a)ttack, (r)etreat, (w)ait"

  for {
    input := askAction(prompt, []string{"a", "r", "w"})

    switch(input) {
    case "a":
      // Attack
      ennemyAvatar.hp = ennemyAvatar.hp - (rollDice(playerAvatar.damageRange) + playerAvatar.damageBonus)
      break
    case "w":
      // Wait
      break
    case "r":
      // Retreat
      return
      break
    }

    if ennemyAvatar.hp <= 0 {
      fmt.Printf("%s is slain...\n", ennemyAvatar.name)
      return
    } else {
      playerAvatar.hp = playerAvatar.hp - (rollDice(ennemyAvatar.damageRange) + ennemyAvatar.damageBonus)
    }

    if playerAvatar.hp <= 0 {
      fmt.Printf("%s is slain...\n", playerAvatar.name)
      return
    }

    showMeleeStatus(playerAvatar, ennemyAvatar)
  }

}

func showMeleeStatus(playerAvatar combatAvatar, ennemyAvatar combatAvatar) {
  showCombatAvatar(playerAvatar)
  fmt.Println("")
  showCombatAvatar(ennemyAvatar)
}

func showCombatAvatar(combatAvatar combatAvatar) {
  fmt.Printf("%s\n", combatAvatar.name)
  fmt.Printf("HP: %d\n", combatAvatar.hp)
  fmt.Printf("AC: %d\n", combatAvatar.ac)
  fmt.Printf("To Hit: %d\n", combatAvatar.tohit)
  fmt.Printf("Damage: 1D%d+%d\n", combatAvatar.damageRange, combatAvatar.damageBonus)
}
