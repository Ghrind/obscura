package main

import "testing"
import "fmt"

func InitTestMod() {
  mod = Mod{}
  mod.AvailableClasses = []AvatarClass{{Name: "warrior"}, {Name: "hunter"}}
}

func TestInitMod(t *testing.T) {
  InitMod()
  for _, monster := range mod.Monsters {
    fmt.Println(monster.Name)
  }
}
