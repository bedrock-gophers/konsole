package app

import (
	"io"
	"net/http"
)

var content []byte

func init() {
	req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/roimee6/roimee6.github.io/main/index.html", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	content, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
}

func (a *App) home(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write(content)
}
