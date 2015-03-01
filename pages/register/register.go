package register

import (
	"fmt"
	"net/http"

	"jam/client"
	"jam/dynamo"
	"jam/madlib"
)

func init() {
	http.HandleFunc("/register/", httpMethod)
}

// boilerplate function to suss out which request we should run
func httpMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET": Get(w, r)
		case "POST": Post(w, r)
	}
}


// Get is the function we run when we send a GET request to the page
func Get(w http.ResponseWriter, r *http.Request) {
	var form dynamo.Obj

	form.Body += renderForm(nil)
	fmt.Fprintf(w, form.Body)
}

// The funky text in the post strings for some of these input elements is
// so we can render in error messages with madlib.
func renderForm(prompts map[string]string) string {
	var form dynamo.Obj

	form.Append("", "INPUT", "required><!-- {username error} --", map[string]string {
		"id" : "name",
		"maxlength" : "50",
		"name" : "name",
		"placeholder" : "username",
		"type" : "text",
		"value" : "<!-- {username value} -->",
	}).BR().Append("", "INPUT", "required><!-- {email error} --", map[string]string {
		"id" : "email",
		"maxlength" : "50",
		"name" : "email",
		"pattern" : "^[A-Za-z\\d]+@[A-Za-z\\d.]+",
		"placeholder" : "email",
		"type" : "text",
	}).BR().Append("", "INPUT", "required><!-- {password error} --", map[string]string {
		"id" : "password",
		"name" : "password",
		"placeholder" : "password",
		"type" : "password",
	}).BR().Append("", "INPUT", "required><!-- {confirm_password error} --", map[string]string {
		"id" : "confirm_password",
                "name" : "confirm_password",
                "placeholder" : "confirm password",
                "type" : "password",
	}).BR().Append("register", "BUTTON", "", map[string]string {
		"formmethod" : "post",
		"id" : "keygen",
		"formaction" : "/register/",
	}).Wrap("FORM", "", nil).Wrap("HTML", "", nil)

	form.Body = madlib.Render(form.Body, prompts)
	return form.Body
}

/*
func renderError(prompt string) string {
	var prompt dynamo.Obj
	prompt.Append(prompt, "SPAN", "", map[string]string {
		"class" : "error",
	})

	return prompt.Body
}
*/

// Get is the function we run when we send a GET request to the page
func Post(w http.ResponseWriter, r *http.Request) {
	var tags map[string]string = make(map[string]string)
	var valid bool = true
	tags["username value"] = r.FormValue("name")
//	email := r.FormValue("email")

	password := r.FormValue("password")
	confirm := r.FormValue("confirm_password")
	tags["password error"], tags["confirm_password error"] = validPassword(password, confirm, &valid)

	if valid {
		e := (client.Register{
			Name: r.FormValue("name"),
			Email: r.FormValue("email"),
			Password: r.FormValue("password"),
		}.Now())

		if e == nil {	
			fmt.Fprintf(w, "SUCCESS")
		} else {
			fmt.Fprintf(w, e.Error())
		}
	} else {
		form := renderForm(tags)
		fmt.Fprintf(w, form)
	}
}

// Uses a pointer to valid rather than returning a valid so we can just
// overwrite the value with false
func validPassword(password, confirm string, valid *bool) (string, string) {
	var perr, cerr string
	perr = "test"

	if password != "confirm" {
		*valid = false
		perr = "password and confirmation do not match!"
	}

	if password == "" {
		*valid = false
		perr = "password required!"
	}

	if confirm == "" {
		*valid = false
		cerr = "confirmation required!"
	}

	return perr, cerr
}
