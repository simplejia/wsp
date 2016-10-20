package controller

import (
	"net/http"
)

// @postfilter("Boss")
func (demo *Demo) Get(w http.ResponseWriter, r *http.Request) {
}
