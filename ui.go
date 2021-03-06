//
// ui.go
//
// Contains all the UI related functions
//
// Naming:
//
// * Use 'display' for partials names
// * Partials don't Flush themselves
// * Use rest-like syntax when possible (NewResource, EditResource)
// * Use 'ask' when getting an input from the user
// * Suffix with 'Screen' when showing a whole new screen
//
package main

import "strconv"
import "fmt"

var currentTerminal Terminal

func initUI(terminal Terminal) {
  currentTerminal = terminal
  currentTerminal.Init()
}

func quitUI() {
  currentTerminal.Close()
  currentTerminal.ExitMessage("Thanks for playing Obscura!")
}

// Return the total number of rows used
func displayItemsList(x int, y int, items []Item, limit int) int {
  var trueLimit int
  if limit > len(items) {
    trueLimit = len(items)
  } else {
    trueLimit = limit
  }
  for i, item := range items[0:trueLimit] {
    currentTerminal.TextAt(x, y + i, fmt.Sprintf("- %s (%d)", item.Name, item.Cost))
  }
  if len(items) > limit {
    currentTerminal.TextAt(x, y + limit, "- ...")
    return limit + 1
  } else {
    return len(items)
  }
}

func ShowAvatarScreen(avatar *Avatar) {
  loop:
  for {
    currentTerminal.Clear()

    currentTerminal.TextAt(0, 0, fmt.Sprintf("%s (%s)", avatar.Name, avatar.Class))
    currentTerminal.TextAt(0, 2, fmt.Sprintf("STR: %d", avatar.Str))
    currentTerminal.TextAt(0, 3, fmt.Sprintf("DEX: %d", avatar.Dex))
    currentTerminal.TextAt(0, 4, fmt.Sprintf("CON: %d", avatar.Con))
    currentTerminal.TextAt(0, 5, fmt.Sprintf("INT: %d", avatar.Int))
    currentTerminal.TextAt(0, 6, fmt.Sprintf("WIL: %d", avatar.Wil))
    currentTerminal.TextAt(0, 7, fmt.Sprintf("PRE: %d", avatar.Pre))

    currentTerminal.TextAt(0, 9, fmt.Sprintf("Loot (%d):", len(avatar.Items)))
    offset := displayItemsList(0, 10, avatar.Items, 10)

    currentTerminal.Flush()

    title := "(e)edit, (f)ight, (q)uit?"
    input := askAction(0, 10 + offset + 1, title, []string{"e", "f", "q"})

    switch input {
    case "e":
      editAvatarScreen(avatar)
    case "f":
      fightScreen(avatar)
    default:
      // "q" and "" (Esc) returns
      break loop
    }
  }
}

func editAvatarScreen(avatar *Avatar) {
  loop:
  for {
    currentTerminal.Clear()

    currentTerminal.TextAt(0, 0, fmt.Sprintf("%s (%s)", avatar.Name, avatar.Class))
    currentTerminal.TextAt(0, 2, fmt.Sprintf("STR: %d", avatar.Str))
    currentTerminal.TextAt(0, 3, fmt.Sprintf("DEX: %d", avatar.Dex))
    currentTerminal.TextAt(0, 4, fmt.Sprintf("CON: %d", avatar.Con))
    currentTerminal.TextAt(0, 5, fmt.Sprintf("INT: %d", avatar.Int))
    currentTerminal.TextAt(0, 6, fmt.Sprintf("WIL: %d", avatar.Wil))
    currentTerminal.TextAt(0, 7, fmt.Sprintf("PRE: %d", avatar.Pre))

    currentTerminal.Flush()

    title := "(r)eroll, (n)ame the character, change the (c)lass of your character, (q)uit?"
    input := askAction(0, 9, title, []string{"r", "n", "c", "q"})

    switch input {
    case "r":
      // Reroll
      rollAvatar(avatar)
    case "n":
      // Name the character
      title := fmt.Sprintf("Give a name to this character (current is '%s')", avatar.Name)
      newName := askString(0, 11, title, "New name?", avatar.Name)
      if newName != "" {
        avatar.Name = newName
      }
    case "c":
      // Change the class
      classNames := make([]string, len(mod.AvailableClasses), len(mod.AvailableClasses))
      for i, class := range mod.AvailableClasses {
        classNames[i] = class.Name
      }
      avatar.Class = askFromList(0, 11, fmt.Sprintf("Choose a class for your character (current is '%s'):", avatar.Class), classNames)
    default:
      // "q" and "" (Esc) returns
      SavegameInterface.Save(*avatar)
      break loop
    }
  }
}

func askAction(x int, y int, prompt string, availableActions []string) string {
  currentTerminal.TextAt(x, y, prompt)
  currentTerminal.Flush()

  for {
    input, err := currentTerminal.WaitKeyPress()
    if err != nil {
      return ""
    }
    for i := 0; i < len(availableActions); i++ {
      if availableActions[i] == input {
        return input
      }
    }
  }

  return ""
}

func askString(x int, y int, title string, prompt string, defaultString string) string {
  currentTerminal.TextAt(x, y, title)
  currentTerminal.TextAt(x, y + 2, prompt)

  currentTerminal.Flush()

  input := string(ShowEditBox(len(prompt) + 2, y + 2, 15, []byte(defaultString)))

  return input
}

func askFromList(x int, y int, title string, list []string) string {
  currentTerminal.TextAt(x, y, title)

  for i:=0; i < len(list); i++ {
    currentTerminal.TextAt(x, y + 2 + i, fmt.Sprintf("%d: %s", i, list[i]))
  }

  prompt := fmt.Sprintf("Your choice (0-%d)?", len(list) - 1)

  currentTerminal.TextAt(x, y + len(list) + 3, prompt)

  currentTerminal.Flush()

  loop:
  for {
    input, err := currentTerminal.WaitKeyPress()
    if err != nil {
      break loop
    }
    choice, err := strconv.Atoi(input)
    if err == nil && choice >= 0 && choice < len(list) {
      return list[choice]
    }
  }

  return ""
}

func fightScreen(avatar *Avatar) {
  ennemy := mod.Monsters[randIndex(len(mod.Monsters))]

  playerAvatar := CombatAvatarFromAvatar(*avatar)
  ennemyAvatar := CombatAvatarFromMonster(ennemy)

  loop:
  for {
    currentTerminal.Clear()
    currentTerminal.TextAt(0, 0, fmt.Sprintf("Melee: %s vs %s\n", playerAvatar.Name, ennemyAvatar.Name))

    displayCombatAvatar(0, 2, playerAvatar)
    displayCombatAvatar(20, 2, ennemyAvatar)

    currentTerminal.Flush()

    prompt := "(a)ttack, (r)etreat, (w)ait?"
    input := askAction(0, 9, prompt, []string{"a", "r", "w"})

    switch(input) {
    case "a":
      // Attack
      ennemyAvatar.Life = ennemyAvatar.Life - RandInt(playerAvatar.DamageMin, playerAvatar.DamageMax)
    case "w":
      // Wait
    default:
      // Retreat / Quit
      showEndScreen(fmt.Sprintf("%s has retreated safely...", playerAvatar.Name))
      break loop
    }

    if ennemyAvatar.Life <= 0 {
      items := PickItems(ennemy.LootMoney)
      avatar.Items = append(avatar.Items, items...)
      SavegameInterface.Save(*avatar)

      currentTerminal.TextAt(0, 11, fmt.Sprintf("%s is slain...", ennemyAvatar.Name))
      currentTerminal.TextAt(0, 13, "The spoils of battle:")
      displayItemsList(0, 14, items, 20)
      currentTerminal.Flush()

      _, _ = currentTerminal.WaitKeyPress()
      break loop
    } else {
      playerAvatar.Life = playerAvatar.Life - RandInt(ennemyAvatar.DamageMin, ennemyAvatar.DamageMax)
    }

    if playerAvatar.Life <= 0 {
      showEndScreen(fmt.Sprintf("%s is slain...", playerAvatar.Name))
      break loop
    }
  }
}

func displayCombatAvatar(x int, y int, combatAvatar CombatAvatar) {
  currentTerminal.TextAt(x, y, combatAvatar.Name)
  currentTerminal.TextAt(x, y + 1, fmt.Sprintf("Life: %d", combatAvatar.Life))
  currentTerminal.TextAt(x, y + 2, fmt.Sprintf("Dodge: %d", combatAvatar.Dodge))
  currentTerminal.TextAt(x, y + 3, fmt.Sprintf("Armor: %d", combatAvatar.Armor))
  currentTerminal.TextAt(x, y + 4, fmt.Sprintf("Attack Rating: %d", combatAvatar.AttackRating))
  currentTerminal.TextAt(x, y + 5, fmt.Sprintf("Damage: %d-%d", combatAvatar.DamageMin, combatAvatar.DamageMax))
}

func showEndScreen(message string) {
  currentTerminal.Clear()

  currentTerminal.TextAt(0, 0, message)
  currentTerminal.TextAt(0, 2, "Press any key to exit")

  currentTerminal.Flush()

  _, _ = currentTerminal.WaitKeyPress()
}

func showErrorScreen(err error) {
  currentTerminal.Clear()

  currentTerminal.TextAt(0, 0, fmt.Sprintf("%s", err))
  currentTerminal.TextAt(0, 2, "Press any key to continue")

  currentTerminal.Flush()

  _, _ = currentTerminal.WaitKeyPress()
}
