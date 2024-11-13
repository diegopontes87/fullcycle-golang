package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name      string
	TotalTime int
}

type Courses []Course

func main() {

	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := tmp.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 23},
		{"Flutter", 90},
	})
	if err != nil {
		panic(err)
	}
}
