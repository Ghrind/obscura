package main

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

type combatAvatar struct {
  name string
  hp int
  ac int
  damageRange int
  damageBonus int
  tohit int
}

func rollAvatar(avatar *avatar) {
  avatar.Str = rollD6() + rollD6() + 6
  avatar.Dex = rollD6() + rollD6() + 6
  avatar.Con = rollD6() + rollD6() + 6
  avatar.Int = rollD6() + rollD6() + 6
  avatar.Wis = rollD6() + rollD6() + 6
  avatar.Cha = rollD6() + rollD6() + 6
}

func modFromStat(stat int) int {
  result := (stat - 10)
  if result % 2 != 0 {
    result = result - 1
  }
  return result / 2
}

func combatAvatarFromAvatar(avatar avatar) combatAvatar {
  combatAvatar := combatAvatar{}
  combatAvatar.name = avatar.Name
  combatAvatar.ac = modFromStat(avatar.Dex) + 10
  combatAvatar.damageRange = 6
  combatAvatar.damageBonus = modFromStat(avatar.Str)
  combatAvatar.hp = 8 + modFromStat(avatar.Con)
  combatAvatar.tohit = modFromStat(avatar.Str) + 1

  return combatAvatar
}
