package contentx

import (
	"html/template"
)

var Tmpl = template.Must(template.ParseGlob("../templates/*"))
//Tmpl is capitallized only capitalized types and functions will be exported(avaliable outside this package) by our package
//contentx.Tmpl This is how we access the template outside this package.
//so we couldn't call contentx.Tmpl if it were lower case.
//template.must is just a utility for template initiallization
//template.parse Glob parses and read all the templates from a directory and loads into our memory whenever the app starts
//it's too usefull as we don't want to read a template from a file everytime we render a template

type Page struct {
	Title string
	Content string
}