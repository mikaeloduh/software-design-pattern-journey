package test

import (
	"net/http"
	"webframework/framework"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var reqData RegisterRequest
	if err := framework.ReadBodyAsObject(r, &reqData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := framework.WriteObjectAsJSON(w, reqData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
