package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/intelfike/comet"
	"github.com/intelfike/shufflechat/io/input"
	"github.com/intelfike/shufflechat/io/output"
	"github.com/intelfike/shufflechat/proc/htmlgen"
	"github.com/intelfike/wshuffle/proc/myshuffle"
)

var port = flag.String("http", ":8888", "HTTP port number.")
var cmt = comet.NewComet("realtimesession")

func init() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/" {
			return
		}
		cmt.Start(w, r)
		s := htmlgen.IndexHTML()
		output.WriteString(w, s)
	})
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		_, text := input.ReadChat(r)
		text = myshuffle.MyShuffle(text, 4, " ")

		name, _ := input.ReadCookie(r, "name")

		// 全てのチャンネルに通知
		cmt.DoneAll(htmlgen.ChatItem(name, text))
	})
	http.HandleFunc("/comet", func(w http.ResponseWriter, r *http.Request) {
		i, err := cmt.Wait(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		switch t := i.(type) {
		case string:
			chatHTML := t
			output.WriteString(w, chatHTML)
		default:
		}
	})
	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		cmt.End(r)
	})
}
func main() {
	fmt.Printf("Start HTTP Server localhost%s\n", *port)
	fmt.Println(http.ListenAndServe(*port, nil))
}
