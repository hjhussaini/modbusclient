package api

import (
	"net/http"

	"github.com/hjhussaini/plc/models"
)

func ResponseError(response http.ResponseWriter, status int, message string) {
	response.WriteHeader(status)
	models.WriteJSON(Result{Message: message}, response)
}
