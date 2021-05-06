package controllers

import (
	"net/http"

	responses "github.com/japhmayor/gocrudapi/api/models/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
