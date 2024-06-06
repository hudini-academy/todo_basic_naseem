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
	mux := pat.New()
	// Routing
	// mux.HandleFunc("/", app.Home)
	mux.Get("/", dynamicMiddleware.ThenFunc(app.Home))
	mux.Post("/addtask", dynamicMiddleware.ThenFunc(app.addTask))
	mux.Post("/deletetask", dynamicMiddleware.ThenFunc(app.deleteTask))
	mux.Post("/updatetask", dynamicMiddleware.ThenFunc(app.UpdateTask))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.ThenFunc(app.logoutUser))
	// CSS file fetching and running
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)

}
