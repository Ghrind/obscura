package main

func initMod() {
  mod = Mod{}
  mod.AvailableClasses = []string{"warrior", "hunter", "sorcerer", "assassin"}
}

var mod Mod

func main() {
  initMod()
  initRandomSeed()

  initUI(new(TermboxTerminal))
  defer quitUI()

  player1, err := load()

  if err != nil {
    showErrorScreen(err)
  }

  if player1.Name == "" {
    player1 = avatar{}
    player1.Name = "unknown"
    player1.Class = "peon"
    rollAvatar(&player1)
  }

  showEditAvatarScreen(&player1)

  save(player1)

  ennemy1 := combatAvatar{name: "Weakling", hp: 4, ac: 11, tohit: 0, damageRange: 4, damageBonus: 0}
  player1Combat := combatAvatarFromAvatar(player1)

  showMeleeScreen(player1Combat, ennemy1)
}
