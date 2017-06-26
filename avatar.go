package main

type Avatar struct {
  Name string
  Class string
  Str int
  Dex int
  Con int
  Int int
  Wil int
  Pre int
  Items []Item
}

type CombatAvatar struct {
  Name string
  Life int
  Dodge int
  DamageMin int
  DamageMax int
  AttackRating int
  Armor int
}

type AvatarClass struct {
  Name string
}

func rollAvatar(avatar *Avatar) {
  avatar.Str = RandInt(1, 6) + RandInt(1, 6) + 6
  avatar.Dex = RandInt(1, 6) + RandInt(1, 6) + 6
  avatar.Con = RandInt(1, 6) + RandInt(1, 6) + 6
  avatar.Int = RandInt(1, 6) + RandInt(1, 6) + 6
  avatar.Wil = RandInt(1, 6) + RandInt(1, 6) + 6
  avatar.Pre = RandInt(1, 6) + RandInt(1, 6) + 6
}

func modFromStat(stat int) int {
  result := (stat - 10)
  if result % 2 != 0 {
    result = result - 1
  }
  return result / 2
}

func CombatAvatarFromAvatar(avatar Avatar) CombatAvatar {
  combatAvatar := CombatAvatar{}
  combatAvatar.Name = avatar.Name
  combatAvatar.Dodge = modFromStat(avatar.Dex) + 10
  combatAvatar.DamageMin = 1 + modFromStat(avatar.Str)
  combatAvatar.DamageMax = 2 + modFromStat(avatar.Str)
  combatAvatar.Life = 6 + modFromStat(avatar.Con)
  combatAvatar.AttackRating = modFromStat(avatar.Str)

  return combatAvatar
}

func CombatAvatarFromMonster(monster Monster) CombatAvatar {
  combatAvatar := CombatAvatar{}
  combatAvatar.Name = monster.Name
  combatAvatar.Life = monster.Life
  combatAvatar.DamageMin = monster.DamageMin
  combatAvatar.DamageMax = monster.DamageMax
  combatAvatar.Life = monster.Life
  combatAvatar.AttackRating = monster.AttackRating

  return combatAvatar
}
