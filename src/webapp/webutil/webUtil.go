package webutil

import (
	"fmt"
	"net/http"	
	"html/template"
)

func ErrorPage(w http.ResponseWriter, message string) {
    fmt.Fprint(w, "<!DOCTYPE html>")
    fmt.Fprint(w, "<html>")
    
    fmt.Fprint(w, "<body>")
    fmt.Fprintf(w, "<h1>ERROR</h1>\n")
    fmt.Fprintf(w, "<h2>%s</h2>\n", message)
    fmt.Fprint(w, "</body>")

    fmt.Fprint(w, "</html>")
}

func GetParameter(request *http.Request, name string) string {
	
	// r.URL.Query() returns a 'Values' type
	// which is simply a map[string][]string of the QueryString parameters.
	queryValues := request.URL.Query()
	
    // Query()["key"] will return an array of items, 
    // we only want a single item => use the first one
	values, ok := queryValues[name]
    if ok && len(values) > 0 {
	    return values[0]
    }
	return ""
}

func Forward(w http.ResponseWriter, filePath string, data interface{} ) {
	
	// Parse the template file
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
	    panic(err)
	}
	
	// Merge the template with the given data to produce the responce
	err = tmpl.Execute(w, data)
	if err != nil {
	    panic(err)
	}
}

