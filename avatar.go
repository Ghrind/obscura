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
  Items []Item
}

type CombatAvatar struct {
  Name string
  Hp int
  Ac int
  DamageRange int
  DamageBonus int
  Tohit int
}

type AvatarClass struct {
  Name string
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

func combatAvatarFromAvatar(avatar avatar) CombatAvatar {
  combatAvatar := CombatAvatar{}
  combatAvatar.Name = avatar.Name
  combatAvatar.Ac = modFromStat(avatar.Dex) + 10
  combatAvatar.DamageRange = 6
  combatAvatar.DamageBonus = modFromStat(avatar.Str)
  combatAvatar.Hp = 8 + modFromStat(avatar.Con)
  combatAvatar.Tohit = modFromStat(avatar.Str) + 1

  return combatAvatar
}

func CombatAvatarFromMonster(monster Monster) CombatAvatar {
  combatAvatar := CombatAvatar{}
  combatAvatar.Name = monster.Name
  combatAvatar.Ac = monster.Ac
  combatAvatar.DamageRange = monster.DamageRange
  combatAvatar.DamageBonus = monster.DamageBonus
  combatAvatar.Hp = monster.Hp
  combatAvatar.Tohit = monster.Tohit

  return combatAvatar
}
