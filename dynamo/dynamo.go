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
		o.Body += "<" + tag + " "
		for k, v := range attr {
			o.Body += k + `="` + v + `" `		
		}
		o.Body += post + ">"

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

// Appends a newline to Obj's Body
func (o *Obj) Newline() *Obj {
	o.Body += "\n"

	return o
}

func (o *Obj) String() string {
	return o.Body
}
