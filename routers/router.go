package routers

import (
	// "github.com/gorilla/handlers"
	"codecomp-backend/controllers"
	"github.com/gorilla/mux"
	// "io/ioutil"
	"net/http"
)

// Init routers for app
func Init() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/", fs)
	userAuth := r.PathPrefix("/auth").Subrouter()
	userAuth.HandleFunc("/github", controllers.GithubOAuthController)
	return r
}
