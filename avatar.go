package main

import "math/rand"

type avatar struct {
  Name string
  Class string
  Str int
  Dex int
  Con int
  Int int
  Wis int
  Cha int
}

func rollAvatar(avatar avatar) avatar {
  avatar.Str = rand.Intn(6) + rand.Intn(6) + 6 + 2
  avatar.Dex = rand.Intn(6) + rand.Intn(6) + 6 + 2
  avatar.Con = rand.Intn(6) + rand.Intn(6) + 6 + 2
  avatar.Int = rand.Intn(6) + rand.Intn(6) + 6 + 2
  avatar.Wis = rand.Intn(6) + rand.Intn(6) + 6 + 2
  avatar.Cha = rand.Intn(6) + rand.Intn(6) + 6 + 2

  return avatar
}
