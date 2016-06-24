package main

import (
	"io"
	"text/template"

	"github.com/topicai/candy"
)

func main() {
	tmpl := template.Must(template.ParseFiles("./template.txt"))

	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}

	candy.WithCreated("./executed_template.txt", func(w io.Writer) {
		candy.Must(tmpl.Execute(w, sweaters))
	})
}
