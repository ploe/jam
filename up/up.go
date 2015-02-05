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


func main() {
	fmt.Println("jam: comin' up... comin' up... comin' up...")

	http.HandleFunc("/", spewUp)
	http.ListenAndServe(":8080", nil)
}