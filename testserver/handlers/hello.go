package handlers

import (
	"fmt"
	"net/http"

	"hotreload/testserver/services"
	"hotreload/testserver/utils"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	t := services.GetCurrentTime()

	msg := utils.FormatMessage("hotreload server")

	fmt.Fprintf(w, "%s\n", msg)
	fmt.Fprintf(w, "name : %s\n", "gokul")
	fmt.Fprintf(w, "time: %s\n", t)
}
