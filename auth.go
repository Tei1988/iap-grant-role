package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func getAuth(w http.ResponseWriter, r *http.Request) {
	validEmailAddress, err := AuthProvider.Authenticate(r)
	if validEmailAddress == "" {
		log.Println("invalid email address", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	role, err := RoleProvider.FindRole(&validEmailAddress)
	if role == nil {
		log.Println(fmt.Sprintf("%s is not granted any role.", validEmailAddress))
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.Header().Set("X-GRANTED-EMAIL-ADDRESS", string(validEmailAddress))

	for k, v := range role {
		w.Header().Set(k, v)
	}
	log.Println(fmt.Sprintf("%s is granted some role.", validEmailAddress))
}

func AuthRouting() chi.Router {
	r := chi.NewRouter()
	r.Get("/", getAuth)

	return r
}
