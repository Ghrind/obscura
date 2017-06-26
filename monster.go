package main

type Monster struct {
  Name string

  // Combat avatar
  Life int
  Dodge int
  Armor int
  DamageMin int
  DamageMax int
  AttackRating int

  // Other
  LootMoney int
}
