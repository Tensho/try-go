package main

import (
	"fmt"
	"os"
	"strconv"

	"./tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
      t, err := strconv.ParseFloat(arg, 64)
      if err != nil {
          fmt.Fprintf(os.Stderr, "%v\n", err)
          os.Exit(1)
      }
      c := tempconv.Celsius(t)
      f := tempconv.Fahrenheit(t)
      k := tempconv.Kelvin(t)
      fmt.Printf("%s\t%s\t%s\n", c, tempconv.CToF(c), tempconv.CToK(c))
      fmt.Printf("%s\t%s\t%s\n", f, tempconv.FToC(f), tempconv.FToK(f))
      fmt.Printf("%s\t%s\t%s\n", k, tempconv.KToC(k), tempconv.KToF(k))
  }
}
