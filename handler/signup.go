package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Waire214/newjwt/auth"
	"github.com/Waire214/newjwt/misc"
	"github.com/Waire214/newjwt/models"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var reg models.RegistrationData

	err := decodeJSONBody(w, r, &reg)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			misc.LogError(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	userRegistration := auth.HandleUserRegistration(reg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userRegistration)
}
