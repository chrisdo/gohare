package web

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"sort"
	"time"

	"github.com/chrisdo/gohare/storage"
)

//go:embed *
var files embed.FS

type Server struct {
	server        http.Server
	storageReader storage.StorageReader
}

func NewServer(addr string, storageReader storage.StorageReader) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/flights", generalFlightsHandler(storageReader))
	return &Server{http.Server{Addr: addr, Handler: mux}, storageReader}

}

func (s *Server) StartServer() {
	s.server.ListenAndServe()
}

func (s *Server) StopServer() {
	s.server.Shutdown(context.Background())

}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := parse("main.html") //later move parse to global to not parse it with each call
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, nil)
}

func generalFlightsHandler(storage storage.StorageReader) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fid := r.URL.Query().Get("fid")
		if len(fid) == 0 {
			err := flightsHandler(w, getFlightList(storage))
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Going for single flit Handler with fid=%s\n", fid)
			err := singleFlightHandler(w, getSingleFlight(storage, fid))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func getFlightList(storage storage.StorageReader) FlightList {
	tracks := storage.GetAllTracks()
	fmt.Printf("Got %d tracks\n", len(tracks))
	result := make([]FlightDataShort, len(tracks))
	for i, track := range tracks {
		result[i] = FlightDataShort{Modes: track.Modes, Callsign: track.Callsign, SSR: track.SSR}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[j].Modes < result[i].Modes
	})
	fmt.Printf("%v\n", result)
	return FlightList{nil, result}
}

func getSingleFlight(storage storage.StorageReader, fid string) FlightData {
	track, err := storage.GetTrackById(fid)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return FlightData{Error: err}
	}
	fmt.Printf("%v\n", track)
	return FlightData{Source: track.Source, Modes: track.Modes, Callsign: track.Callsign,
		SSR: track.SSR, LastUpdate: track.LastUpdate}
}

func parse(file string) (*template.Template, error) {
	return template.New("layout.html").ParseFS(files, "layout.html", file)
}

type FlightData struct {
	Error      error
	Source     string
	Modes      string
	Callsign   string
	SSR        string
	LastUpdate time.Time
}

func singleFlightHandler(w io.Writer, f FlightData) error {
	tmpl, err := parse("flight.html") //later move parse to global to not parse it with each call
	if err != nil {
		return err
	}
	return tmpl.Execute(w, f)
}

type FlightDataShort struct {
	Modes    string
	Callsign string
	SSR      string
}

type FlightList struct {
	Error   error
	Flights []FlightDataShort
}

func flightsHandler(w io.Writer, f FlightList) error {

	tmpl, err := parse("flights.html") //later move parse to global to not parse it with each call
	if err != nil {
		return err
	}
	return tmpl.Execute(w, f)
}
