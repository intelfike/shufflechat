package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/intelfike/shufflechat/io/input"
	"github.com/intelfike/shufflechat/io/output"
	"github.com/intelfike/shufflechat/proc/htmlgen"
	"github.com/intelfike/shufflechat/proc/user"
	"github.com/intelfike/wshuffle/proc/myshuffle"
)

var port = flag.String("http", ":8888", "HTTP port number.")

var chatChanList = make([]chan string, 0)

func init() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := htmlgen.IndexHTML()
		output.WriteString(w, s)
	})
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		_, text := input.ReadChat(r)
		text = myshuffle.MyShuffle(text, 4, " ")

		name, _ := input.ReadCookie(r, "name")
		user.Add(name)

		// 全てのチャンネルに通知
		fmt.Println(chatChanList)
		for _, v := range chatChanList {
			v <- htmlgen.ChatItem(name, text)
		}
		chatChanList = make([]chan string, 0)
	})
	http.HandleFunc("/comet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		ch := make(chan string)
		chatChanList = append(chatChanList, ch)
		chatHTML := <-ch
		output.WriteString(w, chatHTML)
		time.Sleep(time.Second)
	})
}
func main() {
	fmt.Printf("Start HTTP Server localhost%s\n", *port)
	fmt.Println(http.ListenAndServe(*port, nil))
}
