package main

import (
  "bufio"
  "fmt"
  "os"
)

// {
//   "a" => {
//     "filename1" => 2
//     "filename2" => 3
//   },
//   "b" => {
//     "filename1" => 4
//     "filename2" => 2
//   },
// }

func main() {
  counts := make(map[string]map[string]int)
  files := os.Args[1:]
  if len(files) == 0 {
    countLines(os.Stdin, counts)
  } else {
    for _, arg := range files {
      f, err := os.Open(arg)
      if err != nil {
        fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
        continue
      }
      countLines(f, counts)
      f.Close()
    }
  }
  for line, filename_counts := range counts {
    filenames := " "
    total := 0

    for filename, n := range filename_counts {
      filenames += filename + " "
      total += n
    }

    if total > 1 {
      fmt.Printf("%d %s\t[%s]\n", total, line, filenames)
    }
  }
}

func countLines(f *os.File, counts map[string]map[string]int) {
  input := bufio.NewScanner(f)
  for input.Scan() {

    // Create new sub-map if it doesn't yet exist
    if counts[input.Text()] == nil {
      counts[input.Text()] = make(map[string]int)
    }

    counts[input.Text()][f.Name()]++
  }
  // NOTE: ignoring potential errors from input.Err()
}
