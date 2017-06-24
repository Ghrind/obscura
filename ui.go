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

import "strconv"
import "fmt"
import "github.com/nsf/termbox-go"
import "github.com/fatih/color" // Could be easily replaced by control chars

func initUI() {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }

  termbox.SetInputMode(termbox.InputEsc)
}

func quitUI() {
  termbox.Close()
  color.Green("Thanks for playing Crawler!")
}

func showAvatarEditScreen(avatar *avatar) {
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  tbprint(0, 0, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("%s (%s)", avatar.Name, avatar.Class))
  tbprint(0, 2, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("STR: %d", avatar.Str))
  tbprint(0, 3, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("DEX: %d", avatar.Dex))
  tbprint(0, 4, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("CON: %d", avatar.Con))
  tbprint(0, 5, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("INT: %d", avatar.Int))
  tbprint(0, 6, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("WIS: %d", avatar.Wis))
  tbprint(0, 7, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("CHA: %d", avatar.Cha))

  termbox.Flush()

  title := "(r)eroll, (n)ame the character, change the (c)lass of your character, (q)uit?"

  input := askAction(0, 9, title, []string{"r", "n", "c", "q"})

  switch input {
  case "r":
    // Reroll
    rollAvatar(avatar)
    showAvatarEditScreen(avatar)
  case "n":
    // Name the character
    title := fmt.Sprintf("Give a name to this character (current is '%s')", avatar.Name)
    newName := askString(0, 11, title, "New name?", avatar.Name)
    if newName != "" {
      avatar.Name = newName
    }
    showAvatarEditScreen(avatar)
  case "c":
    // Change the class
    classes := mod.AvailableClasses
    avatar.Class = askFromList(0, 11, fmt.Sprintf("Choose a class for your character (current is '%s'):", avatar.Class), classes)
    showAvatarEditScreen(avatar)

  // "q" and "" (Esc) returns
  }
}

func askAction(x int, y int, prompt string, availableActions []string) string {
  tbprint(x, y, termbox.ColorDefault, termbox.ColorDefault, prompt)
  termbox.Flush()

  for {
    ev := termbox.PollEvent()
    if ev.Key == termbox.KeyEsc {
      return ""
    }
    input := string(ev.Ch)
    for i := 0; i < len(availableActions); i++ {
      if availableActions[i] == input {
        return input
      }
    }
  }

  return ""
}

func askString(x int, y int, title string, prompt string, defaultString string) string {
  tbprint(x, y, termbox.ColorDefault, termbox.ColorDefault, title)
  tbprint(x, y + 2, termbox.ColorDefault, termbox.ColorDefault, prompt)

  termbox.Flush()

  input := string(ShowEditBox(len(prompt) + 2, y + 2, 15, []byte(defaultString)))

  termbox.HideCursor()

  return input
}

func askFromList(x int, y int, title string, list []string) string {
  tbprint(x, y, termbox.ColorDefault, termbox.ColorDefault, title)

  for i:=0; i < len(list); i++ {
    tbprint(x, y + 2 + i, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("%d: %s", i, list[i]))
  }

  prompt := fmt.Sprintf("Your choice (0-%d)?", len(list) - 1)

  tbprint(x, y + len(list) + 3, termbox.ColorDefault, termbox.ColorDefault, prompt)

  termbox.Flush()

  for {
    ev := termbox.PollEvent()
    if ev.Key == termbox.KeyEsc {
      return ""
    }
    choice, err := strconv.Atoi(string(ev.Ch))
    if err == nil && choice >= 0 && choice < len(list) {
      return list[choice]
    }
  }

  return ""
}

func showMeleeScreen(playerAvatar combatAvatar, ennemyAvatar combatAvatar) {
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  tbprint(0, 0, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("Melee: %s vs %s\n", playerAvatar.name, ennemyAvatar.name))

  showCombatAvatar(0, 2, playerAvatar)
  showCombatAvatar(20, 2, ennemyAvatar)

  termbox.Flush()

  prompt := "(a)ttack, (r)etreat, (w)ait?"

  input := askAction(0, 8, prompt, []string{"a", "r", "w"})

  switch(input) {
  case "a":
    // Attack
    ennemyAvatar.hp = ennemyAvatar.hp - (rollDice(playerAvatar.damageRange) + playerAvatar.damageBonus)
  case "w":
    // Wait
  case "r":
    // Retreat
    showEndScreen(fmt.Sprintf("%s has retreated safely..."))
    return
  case "":
    // Quit
    return
  }

  if ennemyAvatar.hp <= 0 {
    showEndScreen(fmt.Sprintf("%s is slain...", ennemyAvatar.name))
    return
  } else {
    playerAvatar.hp = playerAvatar.hp - (rollDice(ennemyAvatar.damageRange) + ennemyAvatar.damageBonus)
  }

  if playerAvatar.hp <= 0 {
    showEndScreen(fmt.Sprintf("%s is slain...", playerAvatar.name))
    return
  }

  showMeleeScreen(playerAvatar, ennemyAvatar)

}

func showCombatAvatar(x int, y int, combatAvatar combatAvatar) {
  tbprint(x, y, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("%s\n", combatAvatar.name))
  tbprint(x, y + 1, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("HP: %d\n", combatAvatar.hp))
  tbprint(x, y + 2, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("AC: %d\n", combatAvatar.ac))
  tbprint(x, y + 3, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("To Hit: %d\n", combatAvatar.tohit))
  tbprint(x, y + 4, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("Damage: 1D%d+%d\n", combatAvatar.damageRange, combatAvatar.damageBonus))
}

func showEndScreen(message string) {
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  tbprint(0, 0, termbox.ColorDefault, termbox.ColorDefault, message)
  tbprint(0, 2, termbox.ColorDefault, termbox.ColorDefault, "Press any key to exit")

  termbox.Flush()

  _ = termbox.PollEvent()
}
