// go run 4.10.go repo:golang/go is:open json decoder

package main

import (
  "fmt"
  "log"
  "os"
  "time"

  "../github"
)

func main() {
  result, err := github.SearchIssues(os.Args[1:])
  if err != nil {
    log.Fatal(err)
  }

  yesterdayIssues := make([]*github.Issue, 0)
  monthIssues := make([]*github.Issue, 0)
  yearIssues := make([]*github.Issue, 0)

  now := time.Now()
  day := now.AddDate(0, 0, -1)
  month := now.AddDate(0, -1, 0)
  year := now.AddDate(-1, 0, 0)

  for _, item := range result.Items {
    if item.CreatedAt.After(day) {
      yesterdayIssues = append(yesterdayIssues, item)
    }
    if item.CreatedAt.After(month) {
      monthIssues = append(monthIssues, item)
    }
    if item.CreatedAt.After(year) {
      yearIssues = append(yearIssues, item)
    }
  }

  fmt.Printf("Yesterday\n")
  for _, item := range yesterdayIssues {
    fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
  }

  fmt.Printf("\nMonth\n")
  for _, item := range monthIssues {
    fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
  }

  fmt.Printf("\nYear\n")
  for _, item := range yearIssues {
    fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
  }
}
