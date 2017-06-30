package main

import "github.com/Ghrind/termui"
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

// Termui implementation

type TermuiTerminal struct {
}

func (tb *TermuiTerminal) Init() {
  err := termui.Init()
  if err != nil {
    panic(err)
  }
}

func (tb TermuiTerminal) Close() {
  termui.Close()
}
func (tb TermuiTerminal) Flush() {
  //termui.Flush()
}

func (tb TermuiTerminal) Clear() {
  termui.Clear()
}

func (tb TermuiTerminal) ExitMessage(message string) {
  fmt.Println(message)
}

func (tb TermuiTerminal) TextAt(x int, y int, text string) {
  p := termui.NewPar(text)
  p.Width = len(text)
  p.Height = 1
  p.Border = false
  p.X = x
  p.Y = y

  termui.Render(p)
}

func (tb TermuiTerminal) WaitKeyPress() (string, error) {
  var key string
  var err error

  termui.Handle("/sys/kbd/<escape>", func(e termui.Event) {
    key = ""
    err = errors.New("Escape pressed")
    termui.StopLoop()
  })

  termui.Handle("/sys/kbd", func(e termui.Event) {
    key = e.Data.(termui.EvtKbd).KeyStr
    err = nil
    termui.StopLoop()
  })

  termui.Loop()
  return key, err
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
