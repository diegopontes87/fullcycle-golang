package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name      string
	TotalTime int
}

func main() {
	course := Course{Name: "Go", TotalTime: 40}

	tmp := template.New("CourseTemplate")
	tmp, _ = tmp.Parse("Couse: {{.Name}} - TotalTime{{.TotalTime}}")

	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
