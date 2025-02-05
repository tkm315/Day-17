package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	err := os.WriteFile(filename, p.Body, 0600)
	return err
}
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	p := &Page{Title: title, Body: body}
	return p, nil
}

// create , save , load
func main() {
	p1 := &Page{Title: "TKM315.1", Body: []byte("This is 1st edit")}
	p1.save()
	p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))
}
