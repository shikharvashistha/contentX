//html/template package which is used to render the html from a template.
//net/http package to handle https requests.
//read/write to the database and use it to power content management system.
//templates dir will hold the html templates and cmd will contain the executable code ie. main.go file
package main

import (
	"codes/contentx"
	"os"
)

func main() {
	p := &contentx.Page{ //created a reference to an instance of contentx.Page
		Title:   "Hello World",
		Content: "Body of our web page",
	}
	contentx.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
	//os.Stdout is where to write the template to
	//next argument is the template name
	//next argument is the data to be passed to the template
}
