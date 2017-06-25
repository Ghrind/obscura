package main

var mod Mod

func main() {
  InitMod()
  initRandomSeed()
  InitSavegame(new(TempFileSavegame))

  initUI(new(TermboxTerminal))
  defer quitUI()

  player, err := SavegameInterface.Load()

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
