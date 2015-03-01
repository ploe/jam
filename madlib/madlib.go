//madlib is a really dumb templating format I devised while working on mij
//
//It renders over the top of HTML comments with keys nested between the 
//delimtiers '{.*?}'
//
//You give it a bunch of key value pairs and it will find the key in the
//HTML and replace it with the value.
//
//The beauty of doing it this was is that the templates can be designed in
//regular HTML editing tools and all the dynamic stuff stays separate.
package madlib

import (
	"regexp"
	"strings"
)

//Render takes the string we want to alter and a map of key-value pairs
//that we want to interpolate (tags) 
//
//It returns the string with all the dirty bits replaced. It removes any 
//tags that are unused, by magic...
func Render(body string, tags map[string]string) string {
	re := regexp.MustCompile("<!-- {(.*?)} -->")
	found := re.FindAllStringSubmatch(body, -1)
	for _, v := range found {
		body = strings.Replace(body, v[0], tags[v[1]], -1)	
	}

	return body
}
