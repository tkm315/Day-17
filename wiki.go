package main

//hard coded html edit page
import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

// os.WriteFile
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// make filename , read context , return *Page
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}

// extract title , loadPage(title) , display
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// load or create *page , privide editing form
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, `<h1>%s</h1>`+
		`<form action="/save/%s" method="post">`+`
					<textarea name="body">%s</textarea><br>`+`
					<input type="submit" value="Save"></form>`, p.Title, p.Title, p.Body)
}

// create *page , find title , find body from form , save page
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
}
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
