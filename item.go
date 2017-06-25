package main

type Item struct {
  Name string
  Cost int
}

func PickItems(maxCost int) []Item {
  items := []Item{}

  loop:
  for {
    availableItems := []Item{}

    for _, item := range mod.Items {
      // Ignore items with no cost so we wont loop endlessly
      if item.Cost <= maxCost || item.Cost == 0 {
        availableItems = append(availableItems, item)
      }
    }

    if len(availableItems) == 0 {
      break loop
    }

    item := availableItems[randIndex(len(availableItems))]
    items = append(items, item)
    maxCost -= item.Cost
  }

  return items
}
