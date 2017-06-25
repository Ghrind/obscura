package main

import "testing"

func InitTestMod() {
  mod = Mod{}
  mod.AvailableClasses = []AvatarClass{{Name: "warrior"}, {Name: "hunter"}}
}

func TestInitMod(t *testing.T) {
  InitMod()
  if len(mod.Monsters) == 0 {
    t.Error("It seems that the mod is not loaded properly")
  }
}
