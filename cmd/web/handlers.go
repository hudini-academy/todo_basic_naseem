package main

import (
	"html/template"
	"log"
	"naseem/pkg/models"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	// here we are calling the Latest function to fetch the all items from the database
	getAllItem, err := app.todo.Latest()
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//  retrieving the home template
	files := []string{
		"./ui/html/home.page.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// Execute the template and show the slice of allTodo
	// panic("ur system under maintance")
	err = ts.Execute(w, struct {
		Tasks []*models.Todo
		Flash string
	}{
		Tasks: getAllItem,
		Flash: app.session.PopString(r, "flash"),
	})
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) addTask(w http.ResponseWriter, r *http.Request) {
	// inserting db
	//errors := make(map[string]string)

	todoText := r.FormValue("Name")
	if len(strings.TrimSpace(todoText)) != 0 {
		_, err := app.todo.Insert(todoText)
		if err != nil {
			app.errorLog.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
		app.session.Put(r, "flash", "todo inserted successfully created!")
	} else {
		app.session.Put(r, "flash", " item field cant be empty!")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (app *application) deleteTask(w http.ResponseWriter, r *http.Request) {
	// fetch the id from the form
	id, _ := strconv.Atoi(r.FormValue("id"))
	// deleting the item from db
	_, err := app.todo.Delete(id)
	if err != nil {
		log.Println(err)
	}

	app.session.Put(r, "flash", "item  deleted successfully!")

	// redirect the to the home
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) UpdateTask(w http.ResponseWriter, r *http.Request) {
	// fetch the id from the form
	id, _ := strconv.Atoi(r.FormValue("id"))
	// updating the db

	if len(strings.TrimSpace(r.FormValue("updateTask"))) != 0 {
		_, err := app.todo.Update(id, r.FormValue("updateTask"))

		if err != nil {
			log.Println(err)
		}
		app.session.Put(r, "flash", "item updated successfully!")

	} else {
		app.session.Put(r, "flash", "item Field cant be empty!")

	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Display the user signup form...")
	files := []string{
		"./ui/html/signup.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	ts.Execute(w, nil)
}
func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Create a new user...")
	userName := r.FormValue("name")
	userEmail := r.FormValue("email")
	userPassword := r.FormValue("password")
	err := app.users.Insert(userName, userEmail, userPassword)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Display the user login form...")
	files := []string{
		"./ui/html/login.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	ts.Execute(w, app.session.PopString(r, "flash"))
}
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Authenticate and login the user...")
	userName := r.FormValue("name")
	userPassword := r.FormValue("password")
	isUser, err := app.users.Authenticate(userName, userPassword)
	log.Print(isUser)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	if isUser {
		app.session.Put(r, "Authenticated", true)
		app.session.Put(r, "flash", " Login successfully!")
		http.Redirect(w, r, "/", http.StatusSeeOther)

	} else {
		app.session.Put(r, "flash", " Login Failed!")
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		app.session.Put(r, "Authenticated", true)

	}

}
func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Logout the user...")
	app.session.Put(r, "Authenticated", false)
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

}
