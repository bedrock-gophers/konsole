package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func (a *App) assets(w http.ResponseWriter, r *http.Request) {
	filename := mux.Vars(r)["filename"]
	if len(filename) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b, err := os.ReadFile("./assets/" + filename)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, _ = w.Write(b)
}
