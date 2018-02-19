package chat

import (
	"bytes"
	"fmt"
	"html/template"
)

type Chat struct {
	ID       int
	UserName string
	Text     string
}

func (c *Chat) GetHTML() string {
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
