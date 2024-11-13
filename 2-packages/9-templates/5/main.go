package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name      string
	TotalTime int
}

type Courses []Course

func main() {

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	tmp := template.Must(template.New("content.html").ParseFiles(templates...))
	err := tmp.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 23},
		{"Flutter", 90},
	})
	if err != nil {
		panic(err)
	}
}
