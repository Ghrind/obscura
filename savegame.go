package main

import "fmt"
import "encoding/json"
import "io/ioutil"
import "errors"

var SavegameInterface Savegame = VoidSavegame{}

type Savegame interface {
  Save(avatar Avatar)
  Load() (Avatar, error)
}

func InitSavegame(savegame Savegame) {
  SavegameInterface = savegame
}

// Temporary file implementation
type TempFileSavegame struct {}

func (tf TempFileSavegame) Save(avatar Avatar) {
  content, _ := json.Marshal(avatar)
  err := ioutil.WriteFile("/tmp/crawler.json", content, 0644)
  if err != nil {
    fmt.Println(err)
  }
}

func (tf TempFileSavegame) Load() (Avatar, error) {
  avatar := Avatar{}
  content, err := ioutil.ReadFile("/tmp/crawler.json")
  if err != nil {
    return avatar, err
  }
  err = json.Unmarshal(content, &avatar)
  if err != nil {
    return avatar, err
  }
  return avatar, nil
}

// Void implementation
type VoidSavegame struct {}

func (v VoidSavegame) Save(_ Avatar) {
  // Noop
}

func (v VoidSavegame) Load() (Avatar, error) {
  return Avatar{}, errors.New("Trying to load a VoidSavegame")
}

// Testing implementation
type TestingSavegame struct {
  Avatar Avatar
}

func (t *TestingSavegame) Save(avatar Avatar) {
  t.Avatar = avatar
}

func (t TestingSavegame) Load() (Avatar, error) {
  return t.Avatar, nil
}
