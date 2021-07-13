package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Waire214/newjwt/misc"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var loginInfo models.LoginData

	err := decodeJSONBody(w, r, &loginInfo)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			misc.LogError(err)
			http.Error(w, mr.msg, mr.status)
		} else {
			misc.LogError(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	responseBody := auth.HandleUserLogin(loginInfo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseBody)
}
