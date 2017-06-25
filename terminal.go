package main

import "github.com/nsf/termbox-go"
import "fmt"
import "errors"

type Terminal interface {
  Init() // Init and set input mode
  TextAt(x int, y int, text string)
  WaitKeyPress() (string, error)
  Clear()
  Close()
  Flush()
  ExitMessage(message string)
}

// Termbox implementation

type TermboxTerminal struct {
  proxy func() termbox.Event
}

func (tb *TermboxTerminal) Init() {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }

  termbox.SetInputMode(termbox.InputEsc)

  tb.proxy = termbox.PollEvent
}

func (tb TermboxTerminal) Close() {
  termbox.Close()
}
func (tb TermboxTerminal) Flush() {
  termbox.Flush()
}

func (tb TermboxTerminal) Clear() {
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  termbox.HideCursor()
}

func (tb TermboxTerminal) ExitMessage(message string) {
  fmt.Println(message)
}

func (tb TermboxTerminal) TextAt(x int, y int, text string) {
  tbprint(x, y, termbox.ColorDefault, termbox.ColorDefault, text)
}

func (tb TermboxTerminal) WaitKeyPress() (string, error) {
  ev := tb.proxy()
  switch ev.Type {
  case termbox.EventKey:
    if ev.Key == termbox.KeyEsc {
      return "", errors.New("Escape pressed")
    }

  case termbox.EventError:
    panic(ev.Err)
  }

  return string(ev.Ch), nil
}

// Implementation for automated tests purpose

type TestTerminal struct {
  Content [][]byte
  inputSequence []string
  inputSequenceIndex int
}

func (tt TestTerminal) Init() {
  // Noop
}

func (tt TestTerminal) Close() {
  // Noop
}

func (tt TestTerminal) Flush() {
  // Noop
}

func (tt *TestTerminal) Clear() {
  tt.Content = [][]byte{}
}

func (tt TestTerminal) ExitMessage(message string) {
  tt.Content = [][]byte{[]byte(message)}
}

func (tt *TestTerminal) TextAt(x int, y int, text string) {
  padding := ""
  for i := 0; i < x; i++ {
    padding += " "
  }
  text = padding + text

  // Add new rows if needed
  if len(tt.Content) <= y {
    t := make([][]byte, y+1, (y + 1)*2)
    copy(t, tt.Content)
    tt.Content = t
  }

  tt.Content[y] = []byte(text)
}

func (tt *TestTerminal) WaitKeyPress() (string, error) {
  return tt.nextInputInSequence(), nil
}

func (tt *TestTerminal) ResetInputSequence(sequence []string) {
  tt.inputSequence = sequence
  tt.inputSequenceIndex = 0
}

func (tt *TestTerminal) nextInputInSequence() string {
  if len(tt.inputSequence) <= tt.inputSequenceIndex {
    panic(fmt.Sprintf("Terminal asked for input, but none given at index %d. Given sequence was %s", tt.inputSequenceIndex, tt.inputSequence))
  }
  input := tt.inputSequence[tt.inputSequenceIndex]
  tt.inputSequenceIndex ++
  return input
}
