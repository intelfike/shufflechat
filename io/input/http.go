package input

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/intelfike/nestmap"
)

func ReadChat(r *http.Request) (string, string) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", ""
	}
	defer r.Body.Close()
	i := new(interface{})
	json.Unmarshal(b, i)
	nm := nestmap.New()
	nm.Set(*i)
	name := nm.Child("name").ToString()
	text := nm.Child("text").ToString()
	name, _ = url.QueryUnescape(name)
	text, _ = url.QueryUnescape(text)
	return name, text
}

func ReadCookie(r *http.Request, key string) (string, error) {
	c, err := r.Cookie(key)
	if err != nil {
		return "", errors.New(key + ":cookieにはそのようなキーはありません")
	}
	data, _ := url.QueryUnescape(c.Value)
	return data, nil
}
