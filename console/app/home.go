package app

import (
	"github.com/bedrock-gophers/console/console/template"
	"net/http"
)

var t = template.NewTemplate().WithHTML("frontend/index.html").WithStyle("frontend/style.css").WithScript("frontend/script.js")

func (a *App) home(w http.ResponseWriter, r *http.Request) {
	_ = t.Execute(w)
}
