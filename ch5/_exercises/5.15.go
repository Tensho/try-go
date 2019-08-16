package main

import (
  "fmt"
)

func max(vals ...int) (int, error) {
  if len(vals) == 0 {
    return 0, fmt.Errorf("max: at least one argument should be provided")
  }
  max := vals[0]
  for _, val := range vals[1:] {
    if val > max {
      max = val
    }
  }
  return max, nil
}

func min(vals ...int) (int, error) {
  if len(vals) == 0 {
    return 0, fmt.Errorf("min: at least one argument should be provided")
  }
  min := vals[0]
  for _, val := range vals[1:] {
    if val < min {
      min = val
    }
  }
  return min, nil
}

func main() {
  var values []int

  fmt.Println(max())
  fmt.Println(max(3))
  fmt.Println(max(1, 2, 3, 4))
  fmt.Println(max(-2, 5, 5, 3))
  values = []int{6, 3, 1, 2}
  fmt.Println(max(values...))

  fmt.Println(min())
  fmt.Println(min(3))
  fmt.Println(min(1, 2, 3, 4))
  fmt.Println(min(-2, 5, 5, 3))
  values = []int{6, 3, 1, 2}
  fmt.Println(min(values...))
}
