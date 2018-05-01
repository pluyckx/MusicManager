package server

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func ListenAndServe(port int) error {
	var err error

	tmpl, err = template.New("page").ParseFiles("./html/add_track.html")

	if err != nil {
		return err
	}

	http.HandleFunc("/", handle)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handle(response http.ResponseWriter, request *http.Request) {
	var err error

	response.WriteHeader(200)

	fmt.Println(request.RequestURI)

	data := struct {
		DataArtists  string
		DataLabels   string
		DataTitles   string
		DataReleases string
	}{
		DataArtists:  "<option value=\"Artist1\">\n<option value=\"Artist2\">",
		DataLabels:   "<option value=\"Label1\">\n<option value=\"Label2\">",
		DataTitles:   "<option value=\"Title1\">\n<option value=\"Title2\">",
		DataReleases: "<option value=\"Release1\">\n<option value=\"Release2\">"}

	err = tmpl.Templates()[1].Execute(response, data)

	if err != nil {
		panic(err)
	}
}
