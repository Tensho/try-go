// curl http://localhost:8000/?terms=repo:golang/go&terms=is:open&terms=json&terms=decoder

package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"

	"../github"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>

{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
  	terms := r.URL.Query()["terms"]

  	fmt.Println(terms)

		result, err := github.SearchIssues(terms)
    if err != nil {
			log.Fatal(err)
		}

		if err := issueList.Execute(w, result); err != nil {
			log.Fatal(err)
		}
  })

  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
