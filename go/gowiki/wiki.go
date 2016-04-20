// wiki.go
//
// A wiki class for loading and saving Page data
// structures.
// (c) 2016 zubernetes

package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page")}
	p1.save()
	p2,_ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	
	return &Page{Title: title, Body: body}, nil
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
