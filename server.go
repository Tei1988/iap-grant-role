package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tei1988/iap-grant-role/common"
	"github.com/tei1988/iap-grant-role/provider/auth"
	"github.com/tei1988/iap-grant-role/provider/role"
)

var AuthProvider auth.IAuthProvider
var RoleProvider role.IRoleProvider

func init() {
	config := common.ConfigFactory("config/config.yaml")
	AuthProvider = auth.AuthProviderFactory(config.AuthProvider)
	RoleProvider = role.RoleProviderFactory(config.RoleProvider)
}

func main() {
	root := chi.NewRouter()
	root.Use(middleware.Logger)
	root.Mount("/auth", AuthRouting())
	http.ListenAndServe(":3000", root)
}
