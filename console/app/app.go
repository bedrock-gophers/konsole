package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	endpoint string
	router   *mux.Router
}

func New(endpoint string) *App {
	a := &App{
		endpoint: endpoint,
		router:   mux.NewRouter(),
	}
	return a
}

func (a *App) route() {
	a.router.HandleFunc("/assets/{filename}", a.assets)
	a.router.HandleFunc(a.endpoint, a.home)
}

func (a *App) ListenAndServe(addr string) error {
	a.route()
	return http.ListenAndServe(addr, a.router)
}

func (a *App) ListenAndServeTLS(addr string, certFile, keyFile string) error {
	a.route()
	return http.ListenAndServeTLS(addr, certFile, keyFile, a.router)
}
