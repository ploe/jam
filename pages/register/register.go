package register

import (
	"fmt"
	"net/http"
	"jam/dynamo"
)

func init() {
	http.HandleFunc("/register/", httpMethod)
}

func httpMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET": Get(w, r)
		case "POST": Post(w, r)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
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


func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pages/register => POST")
}
