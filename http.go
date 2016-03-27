package intro
import (
	"net/http"
	"html/template"
	"log"
)

const (
	page = `
		<h1>
			Hello
			{{if .name}}
				{{.name}}
			{{else}}
				Anonymous
			{{end}}
		</h1>
	`
)

func RunServer() {
	templ, err := template.New("hello").Parse(page)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/hello", func (w http.ResponseWriter, req *http.Request) {
		templ.Execute(w, req.URL.Query())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
