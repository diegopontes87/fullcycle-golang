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

	tmp := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Total Time {{.TotalTime}}"))
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
