//
// mod.go
//
// A mod contains all the metadata of the game (ex: which classes are available)
package main

import "os"
import "github.com/gocarina/gocsv"
import "path"

const ModfileDir = "mods/hive" // NOTE: May not be compatible with Windows
const ModfileSuffix = ".csv"

type Mod struct {
  AvailableClasses []AvatarClass
  Monsters []Monster
  Items []Item
}

func InitMod() {
  mod = Mod{}

  loadModFile("classes", &mod.AvailableClasses)
  loadModFile("monsters", &mod.Monsters)
  loadModFile("items", &mod.Items)
}

func ModfilePath(filename string) string {
  return path.Join(ModfileDir, filename + ModfileSuffix)
}

func loadModFile(filename string, out interface{}) {
  file, err := os.Open(ModfilePath(filename))
  defer file.Close()

  if err != nil {
    panic(err)
  }

  err = gocsv.UnmarshalFile(file, out)

  if err != nil {
    panic(err)
  }
}
