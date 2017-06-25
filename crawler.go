package main

var mod Mod

func main() {

  InitMod()
  initRandomSeed()

  initUI(new(TermboxTerminal))
  defer quitUI()

  player, err := load()

  if err != nil {
    showErrorScreen(err)
  }

  if player.Name == "" {
    player = Avatar{}
    player.Name = "unknown"
    player.Class = "peon"
    rollAvatar(&player)
  }

  ShowAvatarScreen(&player)
}
