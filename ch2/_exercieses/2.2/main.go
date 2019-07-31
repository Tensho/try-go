package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"

  "./lengthconv"
  "./tempconv"
  "./weightconv"
)

func main() {
  args := os.Args[1:]

  var numbers []string

  if len(args) == 0 {
    reader := bufio.NewReader(os.Stdin)
    text, err := reader.ReadString('\n')
    if err != nil {
      fmt.Fprintf(os.Stderr, "%v\n", err)
      os.Exit(1)
    }
    numbers = strings.Split(strings.Trim(text, "\n"), " ")
  } else {
    numbers = args
  }

  for _, number := range numbers {
    v, err := strconv.ParseFloat(number, 64)
    if err != nil {
      fmt.Fprintf(os.Stderr, "cf: %v\n", err)
      os.Exit(1)
    }

    convert_temp(v)
    convert_length(v)
    convert_weight(v)

    fmt.Println("--------------------------------------------------------------------------------")
  }
}

func convert_temp(v float64) {
  fmt.Println("Temperature")
  c := tempconv.Celsius(v)
  f := tempconv.Fahrenheit(v)
  k := tempconv.Kelvin(v)
  fmt.Printf("%s\t%s\t%s\n", c, tempconv.CToF(c), tempconv.CToK(c))
  fmt.Printf("%s\t%s\t%s\n", f, tempconv.FToC(f), tempconv.FToK(f))
  fmt.Printf("%s\t%s\t%s\n", k, tempconv.KToC(k), tempconv.KToF(k))
}

func convert_length(v float64) {
  fmt.Println("Length")
  m := lengthconv.Meter(v)
  f := lengthconv.Feet(v)
  fmt.Printf("%s\t%s\n", m, lengthconv.MToF(m))
  fmt.Printf("%s\t%s\n", f, lengthconv.FToM(f))
}

func convert_weight(v float64) {
  fmt.Println("Weight")
  k := weightconv.Kilogram(v)
  p := weightconv.Pound(v)
  fmt.Printf("%s\t%s\n", k, weightconv.KToP(k))
  fmt.Printf("%s\t%s\n", p, weightconv.PToK(p))
}
