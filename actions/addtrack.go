package actions

import (
	"html/template"
	"net/http"
)

const (
	addTrackArtist  = "artist"
	addTrackTitle   = "title"
	addTrackLabel   = "label"
	addTrackRelease = "release"
)

// AddTrackHandler handles all add track requests
type AddTrackHandler struct {
	view       addTrackViewHandler
	controller addTrackAction
}

// addTrackView generates a view to add a new track
type addTrackViewHandler struct {
	template *template.Template
}

// addTrackAction accepts input and add a new track
type addTrackAction struct {
	template *template.Template
}

type addTrackViewData struct {
	Artists  []string
	Labels   []string
	Titles   []string
	Releases []string
}

type addTrackAddData struct {
	Artist  string
	Label   string
	Title   string
	Release string
}

// NewAddTrackHandler creates a new Model View Controll handler to serve add track requests
func NewAddTrackHandler() *AddTrackHandler {
	ath := AddTrackHandler{}
	ath.view.template = template.Must(template.ParseFiles("./html/add_track_input.html"))
	ath.controller.template = template.Must(template.ParseFiles("./html/add_track_add.html"))

	return &ath
}

func (ath *AddTrackHandler) ServeHTTP(out http.ResponseWriter, request *http.Request) {
	action := request.URL.Query().Get("action")
	if action == actionAdd {
		ath.controller.serveHTTP(out, request)
	} else {
		ath.view.serveHTTP(out, request)
	}
}

func (atv *addTrackViewHandler) serveHTTP(out http.ResponseWriter, request *http.Request) {
	data := addTrackViewData{}

	data.Artists = append(data.Artists, "Headhunterz", "Digital Punk", "The Prophet", "Alpha Twins")
	data.Labels = append(data.Labels, "Scantraxx", "Funsion Records", "Italian Hardstyle")
	data.Titles = append(data.Titles, "Unleashed", "United As One", "Unborn")
	data.Releases = append(data.Releases, "Digital Sampler", "Unleashed", "Unborn", "Another Release")

	err := atv.template.Execute(out, data)

	if err != nil {
		panic(err)
	}
}

func (ath *addTrackAction) serveHTTP(out http.ResponseWriter, request *http.Request) {
	data := addTrackAddData{}

	err := request.ParseForm()

	if err != nil {
		panic(err)
	}

	data.Artist = request.Form.Get(addTrackArtist)
	data.Label = request.Form.Get(addTrackLabel)
	data.Release = request.Form.Get(addTrackRelease)
	data.Title = request.Form.Get(addTrackTitle)

	err = ath.template.Execute(out, data)

	if err != nil {
		panic(err)
	}
}
