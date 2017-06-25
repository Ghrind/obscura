package main

import "testing"
import "fmt"

var testTerminal = new(TestTerminal) // Don't know if it is correct, because it will be overidden when calling initTestTerminal()

func initTestTerminal() {
  testTerminal = new(TestTerminal)
  initUI(testTerminal)
}

func terminalContent() string {
  content := ""
  for _, bytes := range testTerminal.Content {
    str := string(bytes)
    content += fmt.Sprintf("%s\n", str)
  }
  return content[:len(content)-1] // Remove trailing newline
}

func expectOutputEquals(t *testing.T, expectedContent string) {
  currentContent := terminalContent()
  if currentContent != expectedContent {
    t.Error(fmt.Sprintf("Expected terminal output to equal \"%s\", got \"%s\"", expectedContent, currentContent))
  }
}

func TestShowCombatAvatar(t *testing.T) {
  initTestTerminal()
  combatAvatar := CombatAvatar{Name: "Foobar", Ac: 10, Hp: 20, Tohit: 2, DamageRange: 6, DamageBonus: 4}
  displayCombatAvatar(0, 0, combatAvatar)

  expectedOutput := "Foobar\n" +
  "HP: 20\n" +
  "AC: 10\n" +
  "To Hit: 2\n" +
  "Damage: 1D6+4"

  expectOutputEquals(t, expectedOutput)
}

func TestShowEditAvatarScreen(t *testing.T) {
  initTestTerminal()
  testTerminal.ResetInputSequence([]string{"q"})

  avatar := avatar{Name: "Morgoth", Class: "Stalker", Str: 1, Dex: 2, Con: 3, Cha: 4, Wis: 5, Int: 6}

  editAvatarScreen(&avatar)

  expectedOutput := "Morgoth (Stalker)\n\n" +
  "STR: 1\n" +
  "DEX: 2\n" +
  "CON: 3\n" +
  "INT: 6\n" +
  "WIS: 5\n" +
  "CHA: 4\n\n" +
  "(r)eroll, (n)ame the character, change the (c)lass of your character, (q)uit?"

  expectOutputEquals(t, expectedOutput)
}

func TestChangeAvatarClass(t *testing.T) {
  initTestTerminal()
  InitTestMod()
  testTerminal.ResetInputSequence([]string{"c", "0", "q"})

  avatar := avatar{Name: "Morgoth", Class: "Stalker", Str: 1, Dex: 2, Con: 3, Cha: 4, Wis: 5, Int: 6}

  editAvatarScreen(&avatar)

  expectedOutput := "Morgoth (warrior)\n\n" +
  "STR: 1\n" +
  "DEX: 2\n" +
  "CON: 3\n" +
  "INT: 6\n" +
  "WIS: 5\n" +
  "CHA: 4\n\n" +
  "(r)eroll, (n)ame the character, change the (c)lass of your character, (q)uit?"

  expectOutputEquals(t, expectedOutput)
}

func TestDisplayItemsListWithNoLimit(t *testing.T) {
  initTestTerminal()
  items := []Item{{Name: "Foo", Cost: 1}, {Name: "Bar", Cost: 2}, {Name: "Baz", Cost: 3}}

  expectedOffset := 3
  offset := displayItemsList(1, 2, items, 100)
  if offset != expectedOffset {
    t.Error(fmt.Sprintf("Expected an offset of %d, got %d", expectedOffset, offset))
  }

  expectedOutput := "\n\n" +
    " - Foo (1)\n" +
    " - Bar (2)\n" +
    " - Baz (3)"

  expectOutputEquals(t, expectedOutput)
}

func TestDisplayItemsListWithLimit(t *testing.T) {
  initTestTerminal()
  items := []Item{{Name: "Foo", Cost: 1}, {Name: "Bar", Cost: 2}, {Name: "Baz", Cost: 3}}

  expectedOffset := 2
  offset := displayItemsList(1, 2, items, 1)
  if offset != expectedOffset {
    t.Error(fmt.Sprintf("Expected an offset of %d, got %d", expectedOffset, offset))
  }

  expectedOutput := "\n\n" +
    " - Foo (1)\n" +
    " - ..."

  expectOutputEquals(t, expectedOutput)
}
