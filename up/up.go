package main

import (
	"fmt"
	"net/http"
	"jam/dynamo"
)

func spewUp(w http.ResponseWriter, r *http.Request)  {
	var href dynamo.Obj
	href.Append("press here", "A", "", map[string]string {
		"href" : "/profile",
		"id" : "profilebutton",
		"class" : "specialneeds",
	}).Newline().Append("or here", "SPAN", "visible", map[string]string {
		"href" : "/testurl",
	}).Wrap("DIV", "ummm",  map[string]string {
		"class" : "nom",
	})

	fmt.Println("Some bastard just requested something...")
	fmt.Fprintf(w, href.Body)
}

func register(w http.ResponseWriter, r *http.Request) {
	var form dynamo.Obj

	form.Append("", "INPUT", "required", map[string]string {
		"name" : "username",
		"id" : "username",
		"maxlength" : "50",
		"type" : "text",
	}).BR().Append("", "INPUT", "required", map[string]string {
		"name" : "password",
		"id" : "password",
		"type" : "password",
	}).BR().Append("press me!", "BUTTON", "", map[string]string {
		"formmethod" : "post",
		"id" : "keygen",
		"formaction" : "/method/",
	}).Radios("gender", "both", map[string]string {
		"male" : "m",
		"female" : "f",
		"neither" : "n",
		"neuter" : "x",
		"both" : "b",
	}).Wrap("FORM", "", nil).Wrap("HTML", "", nil)
	fmt.Fprintf(w, form.Body)
}

func httpMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			fmt.Fprintf(w, "GET")
		case "POST":
			fmt.Fprintf(w, "POST: " + r.FormValue("password"))
	}
}

func main() {
	fmt.Println("jam: comin' up... comin' up... comin' up...")

	http.HandleFunc("/", spewUp)
	http.HandleFunc("/method/", httpMethod)
	http.HandleFunc("/register/", register)
	http.ListenAndServe(":8080", nil)
}
