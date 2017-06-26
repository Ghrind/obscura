# Obscura

## Build & Run

    go build && ./obscura

## Resources to lookup to

* https://github.com/JoelOtter/termloop a game engine (game loop and game levels)

## Game rules

* Simple base mechanic with various special rules
* Combat oriented stats



* AttackRating: ability to touch the target in order to inflict damage
* Dodge: Ability to avoid being hit
* DamageMin: Minimum damage inflicted when hitting a target (may change when using another weapon)
* DamageMax: Maximum damage inflicted when hitting a target (may change when using another weapon)
* Armor: DamageResistance that comes from physical protection
* Life: Amount of damage that can be sustained before being out of action
* Morale: Amount of psychic damage that can be sustained before fleeing. Psychic damage is inflicted by long fights and some gruesome attacks

Avatars have base characteristics:

* Strength: Adds a bonus to physical melee damages, allows to bear more weight
* Constitution: Adds a bonus to Life
* Dexterity: Adds a bonus to AttackRating and Dodge
* Willpower: Adds a bonus to Morale
* Intelligence: Ability to use sophisticated technologies
* Presence: Amount of psychic damage inflicted at each combat cycle
