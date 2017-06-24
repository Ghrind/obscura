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

type termboxImpl struct {
  // TODO: Add version
  proxy func() termbox.Event
}

func (tb *termboxImpl) Init() {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }

  termbox.SetInputMode(termbox.InputEsc)

  tb.proxy = termbox.PollEvent
}

func (tb termboxImpl) Close() {
  termbox.Close()
}
func (tb termboxImpl) Flush() {
  termbox.Flush()
}

func (tb termboxImpl) Clear() {
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  termbox.HideCursor()
}

func (tb termboxImpl) ExitMessage(message string) {
  fmt.Println(message)
}

func (tb termboxImpl) TextAt(x int, y int, text string) {
  tbprint(x, y, termbox.ColorDefault, termbox.ColorDefault, text)
}

func (tb termboxImpl) WaitKeyPress() (string, error) {
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
