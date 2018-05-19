package main

import (
	"os"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"github.com/odusanya18/modo-embeds/formats"
)

const (
	projectID = "my-project-id"
)

var (
	iframe = template.Must(template.ParseFiles("views/iframe.html"))
)

func main() {
	registerHandlers()
	appengine.Main()
}

func registerHandlers(){
	r := mux.NewRouter()
	r.Methods("GET").Path("/video/{id}").HandlerFunc(iframeHandler)
	r.Methods("GET").Path("/_ah/health").HandlerFunc(healthcheckHandler)
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}

func iframeHandler(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	key := datastore.NewKey(ctx, "video", mux.Vars(r)["id"] , 0, nil) 
	vid := new(formats.Video)
    if err := datastore.Get(ctx, key, &vid); err != nil {
		http.Error(w, err.Error(), 404)
        return
	}
    if err := iframe.Execute(w, vid); err != nil {
        http.Error(w, err.Error(), 500)
    }
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
