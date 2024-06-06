package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders, app.logResponse)
	// mux := http.NewServeMux()
	dynamicMiddleware := alice.New(app.session.Enable)
	// this will check the authenication middleware 
	dynamicMiddleware2 := alice.New(app.session.Enable, app.authenticateMiddleware)
	mux := pat.New()
	// Routing
	// mux.HandleFunc("/", app.Home)
	mux.Get("/", dynamicMiddleware2.ThenFunc(app.Home))
	mux.Post("/addtask", dynamicMiddleware2.ThenFunc(app.addTask))
	mux.Post("/deletetask", dynamicMiddleware2.ThenFunc(app.deleteTask))
	mux.Post("/updatetask", dynamicMiddleware2.ThenFunc(app.UpdateTask))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware2.ThenFunc(app.logoutUser))
	// CSS file fetching and running
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)

}
