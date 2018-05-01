package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/pluyckx/MusicManager/actions"
)

var tmpl *template.Template

// ListenAndServe starts a http server, this function will block
func ListenAndServe(port int) error {
	http.Handle("/add_track", actions.NewAddTrackHandler())

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
