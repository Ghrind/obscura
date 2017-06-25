package main

var mod Mod

func main() {

  InitMod()
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

  showAvatarScreen(&player1)

  save(player1)

  ennemy1 := mod.Monsters[randIndex(len(mod.Monsters))]

  player1Combat := combatAvatarFromAvatar(player1)
  ennemy1Combat := CombatAvatarFromMonster(ennemy1)

  showMeleeScreen(player1Combat, ennemy1Combat)

  player1.Items = append(player1.Items, PickItems(ennemy1.LootMoney)...)
  save(player1)
}
