package handler

import (
	"dreamai/view/home"
	"net/http"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) {
	home.Index().Render(w)
}
