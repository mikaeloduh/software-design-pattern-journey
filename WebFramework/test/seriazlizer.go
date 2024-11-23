package test

import (
	"net/http"
	"webframework/framework"
)

type RegisterRequest struct {
	Username string `json:"username" xml:"username"`
	Email    string `json:"email" xml:"email"`
	Password string `json:"password" xml:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var reqData RegisterRequest
	if err := framework.ReadBodyAsObject(r, &reqData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accept := r.Header.Get("Accept")
	if accept == "application/xml" {
		if err := framework.WriteObjectAsXML(w, reqData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := framework.WriteObjectAsJSON(w, reqData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
