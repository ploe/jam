// dynamo stands for "dynamic objects" - this package is for rendering
// dynamic HTML elements. It provides a type that lets you slot together
// chunks of HTML using maps for the most part.
//
// The package also provides a few helper functions in the same spirit!
package dynamo

// A dynamo Obj is a type used to build HTML elements using maps.
// Doing it this way just seems tight and clean, and easily editable.
// It's a little Perl-ish I admit, but I feel you need that sort of 
// hackability when you're dabbling in HTML.
type Obj struct {
	Body string
}

// Appends an HTML element to the dynamo Obj's body. content is the text 
// you want between the HTML tags. tag is the type of HTML tag. post is
// the boolean attributes inside the angle brackets. attr is a map of 
// strings that gets used for the element's attributes their key being the
// attribute name and the value being the attribute value.
func (o *Obj) Append(content, tag, post string, attr map[string]string) *Obj {
	if tag != "" {
		o.Body += "<" + tag
		for k, v := range attr {
			o.Body += " " + k + `="` + v + `"`		
		}
		if post != "" { o.Body += " " + post }
		o.Body += ">"


		if content != "" {
			o.Body += content +	`</` + tag + `>`
		} 
	}

	return o
}

// Like Append, except it destroys everything that is currently in the 
// Obj's Body
func (o *Obj) Write(content, tag, post string, attr map[string]string) *Obj {
	o.Body = ""
	o.Append(content, tag, post, attr)
	return o
}

// Like Append except it wraps the current Body in the HTML element we're 
// building.
func (o *Obj) Wrap(tag, post string, attr map[string]string) *Obj {
	body := "\n" + o.Body + "\n"
	o.Write(body, tag, post, attr)

	return o
}

// Generates a load of radio buttons in the body of the Obj. name is the
// name of the http param, def is the default checked and attr is a bunch
// of key value pairs; The key being the label and the value being what we
// set the param to. The id is the name and key separated by an underscore.
func (o *Obj) Radios(name, def string, attr map[string]string) *Obj {
	for k, v := range attr {
		id := name + "_" + k
		var checked string
		if def == k { checked = "checked" }
		o.Append("", "INPUT", checked, map[string]string {
			"id" : id,
			"name" : name,
			"type" : "radio",
			"value" : v,
		}).Append(k, "LABEL", "", map[string]string {
			"for" : id,
		}).Newline()
	}
	
	return o
}

// Appends a newline to Obj's Body
func (o *Obj) Newline() *Obj {
	o.Body += "\n"

	return o
}

func (o *Obj) BR() *Obj {
        o.Body += "<BR>\n"

	return o
}

func (o *Obj) String() string {
	return o.Body
}
