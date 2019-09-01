package main

import (
  "fmt"
  "os"
  "sort"
  "text/tabwriter"
  "time"
)

type Track struct {
  Title  string
  Artist string
  Album  string
  Year   int
  Length time.Duration
}

var tracks = []*Track{
  {"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
  {"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
  {"Go", "Bihan", "Another World", 2012, length("4m17s")},
  {"Go", "Moby", "Moby", 1992, length("3m37s")},
  {"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
  d, err := time.ParseDuration(s)
  if err != nil {
    panic(s)
  }
  return d
}

func printTracks(mt MultiTireTracks) {
  const format = "%v\t%v\t%v\t%v\t%v\t\n"
  tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
  fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
  fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
  for _, t := range mt.tracks {
    fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
  }
  tw.Flush() // calculate column widths and print table
}

type MultiTireTracks struct {
  tracks []*Track
  tires []string
}

func (x MultiTireTracks) Len() int           { return len(x.tracks) }
func (x MultiTireTracks) Swap(i, j int)      { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }
func (x MultiTireTracks) Less(i, j int) bool {
  fmt.Println("Less: %v", x.tires)
  for _, tire := range x.tires {
    if tire == "Title" && x.tracks[i].Title != x.tracks[j].Title {
      return x.tracks[i].Title < x.tracks[j].Title
    }

    if tire == "Artist" && x.tracks[i].Artist != x.tracks[j].Artist {
      return x.tracks[i].Artist < x.tracks[j].Artist
    }

    if tire == "Album" && x.tracks[i].Album != x.tracks[j].Album {
      return x.tracks[i].Album < x.tracks[j].Album
    }

    if tire == "Year" && x.tracks[i].Year != x.tracks[j].Year {
      return x.tracks[i].Year < x.tracks[j].Year
    }

    if tire == "Length" && x.tracks[i].Length != x.tracks[j].Length {
      return x.tracks[i].Length < x.tracks[j].Length
    }
  }
  return false
}

func (x *MultiTireTracks) SetTire(s string) {
  x.tires = append([]string{s}, x.tires...) // prepend
}

// Title       Artist          Album              Year  Length
// -----       ------          -----              ----  ------
// Go          Delilah         From the Roots Up  2012  3m38s
// Go Ahead    Alicia Keys     As I Am            2007  4m36s
// Go          Bihah           Another World      2012  4m17s
// Go          Moby            Moby               1992  3m37s
// Ready 2 Go  Martin Solveig  Smash              2011  4m24s

// Sort by: Title, Year, Artist
// Title       Artist          Album              Year  Length
// -----       ------          -----              ----  ------
// Go          Moby            Moby               1992  3m37s
// Go          Bihah           Another World      2012  4m17s
// Go          Delilah         From the Roots Up  2012  3m38s
// Go Ahead    Alicia Keys     As I Am            2007  4m36s
// Ready 2 Go  Martin Solveig  Smash              2011  4m24s
func main() {
  mt := MultiTireTracks{tracks, []string{}}
  printTracks(mt)

  fmt.Println("\nby Artist:\n")
  mt.SetTire("Artist")
  sort.Sort(mt)
  printTracks(mt)

  fmt.Println("\nby Year, Artist:\n")
  mt.SetTire("Year")
  sort.Sort(mt)
  printTracks(mt)

  fmt.Println("\nby Title, Year, Artist:\n")
  mt.SetTire("Title")
  sort.Sort(mt)
  printTracks(mt)
}
