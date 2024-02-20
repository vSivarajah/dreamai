package handler

import (
	"net/http"

	"github.com/vSivarajah/dreamai/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) {
	home.Index().Render(w)
}
