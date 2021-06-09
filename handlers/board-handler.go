package handlers

import (
	"fmt"
	"net/http"
)

func GetBoardList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "board List")
}
