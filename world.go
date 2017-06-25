package main

import "fmt"
import "encoding/json"
import "io/ioutil"

func save(avatar Avatar) {
  content, _ := json.Marshal(avatar)
  err := ioutil.WriteFile("/tmp/crawler.json", content, 0644)
  if err != nil {
    fmt.Println(err)
  }
}

func load() (Avatar, error) {
  avatar := Avatar{}
  content, err := ioutil.ReadFile("/tmp/crawler.json")
  if err != nil {
    return avatar, err
  }
  err = json.Unmarshal(content, &avatar)
  if err != nil {
    return avatar, err
  }
  return avatar, nil
}
