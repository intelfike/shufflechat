package htmlgen

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/intelfike/shufflechat/io/input/file"
)

func IndexHTML() string {
	s, _ := file.ReadFile("data/index.html")
	return s
}

func ChatItem(name, text string) string {
	tmpl, err := template.ParseFiles("data/chatItem.html")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	noItems := struct {
		Name string
		Text string
	}{
		Name: name,
		Text: text,
	}
	b := new(bytes.Buffer)
	tmpl.Execute(b, noItems)
	return b.String()
}
