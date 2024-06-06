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

	// checking whther the field length is not empty then upadte ,else show message field cant be empty
	if len(strings.TrimSpace(r.FormValue("updateTask"))) != 0 {
		// updating the db
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
	// we are parsing the file
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// here we are executing the file and attaching flash for the session message
	ts.Execute(w, app.session.PopString(r, "flash"))
}
func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Create a new user...")
	// we are fetching the values  from form
	userName := r.FormValue("name")
	userEmail := r.FormValue("email")
	userPassword := r.FormValue("password")
	// inserting users database
	err := app.users.Insert(userName, userEmail, userPassword)
	// if the users is there ,it shows users already exist and redirect to sign up page
	// else show the message sign up successfully and redirecting to home page and also we are passing a key to middleware
	if err != nil {
		app.errorLog.Println(err.Error())
		app.session.Put(r, "flash", "User already exist")
		http.Redirect(w, r, "/user/signup", http.StatusSeeOther)
	} else {
		app.session.Put(r, "Authenticated", true)
		app.session.Put(r, "flash", "sign up successfull!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Display the user login form...")
	files := []string{
		"./ui/html/login.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	// parsing the file pages
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// here we are executing the file and attaching flash for the session message
	ts.Execute(w, app.session.PopString(r, "flash"))
}
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Authenticate and login the user...")
	// fetching values from the form
	userName := r.FormValue("name")
	userPassword := r.FormValue("password")
	// we are authenticating whether the username and password is correct and return true
	isUser, err := app.users.Authenticate(userName, userPassword)
	log.Print(isUser)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	// if the user is true,we pass key to middleware and return true to middleware and also show login successfull message
	if isUser {
		app.session.Put(r, "Authenticated", true)
		app.session.Put(r, "flash", " Login successfully!")
		http.Redirect(w, r, "/", http.StatusSeeOther)

	} else {
		app.session.Put(r, "flash", " name or password is incorrect!")
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		app.session.Put(r, "Authenticated", false)

	}

}
func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Logout the user...")
	// here we are passing authentiction failed and redirect to login page

	app.session.Put(r, "Authenticated", false)
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

}
