package ascii

import (
	"html/template"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmplt/index.html"))
	r.ParseForm()
	banner := r.Form.Get("select")
	input := r.Form.Get("user_input")
	if r.URL.Path != "/" { // handel if url was not valide
		http.Error(w,"page not found",http.StatusNotFound)
		return
	}

	v := Printing(input, banner)

	tmpl.Execute(w, v)
}
