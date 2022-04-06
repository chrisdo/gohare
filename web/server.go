package web

import (
	"context"
	"embed"
	"html/template"
	"io"
	"net/http"
	"sort"
	"time"

	"github.com/chrisdo/gohare/storage"
	"github.com/rs/zerolog/log"
)

//go:embed *
var files embed.FS

var (
	home         = parse("main.html")
	flights      = parse("flights.html")
	singleFlight = parse("flight.html")
)

type Server struct {
	server        http.Server
	storageReader storage.StorageReader
}

type FlightData struct {
	Error      error
	Source     string
	Modes      string
	Callsign   string
	SSR        string
	LastUpdate time.Time
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
	log.Trace().Msgf("call to main from %s", r.RemoteAddr)
	home.Execute(w, nil)
}

func singleFlightHandler(w io.Writer, f FlightData) error {
	return singleFlight.Execute(w, f)
}

func flightsHandler(w io.Writer, f FlightList) error {
	return flights.Execute(w, f)
}

func generalFlightsHandler(storage storage.StorageReader) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fid := r.URL.Query().Get("fid")
		if len(fid) == 0 {
			log.Info().Msg("Handling flight list")
			err := flightsHandler(w, getFlightList(storage))
			if err != nil {
				log.Err(err)
			}
		} else {
			log.Info().Str("ModesId", fid).Msg("Handling single flight")
			err := singleFlightHandler(w, getSingleFlight(storage, fid))
			if err != nil {
				log.Err(err)
			}
		}
	}
}

func getFlightList(storage storage.StorageReader) FlightList {
	tracks := storage.GetAllTracks()
	log.Debug().Msgf("Got %d tracks\n", len(tracks))
	result := make([]FlightDataShort, len(tracks))
	for i, track := range tracks {
		result[i] = FlightDataShort{Modes: track.Modes, Callsign: track.Callsign, SSR: track.SSR}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[j].Modes < result[i].Modes
	})
	log.Debug().Msgf("Tracks: %v\n", result)
	return FlightList{nil, result}
}

func getSingleFlight(storage storage.StorageReader, fid string) FlightData {
	track, err := storage.GetTrackById(fid)
	if err != nil {
		log.Err(err)
		return FlightData{Error: err}
	}
	log.Debug().Str("Modesid", fid).Msgf("%v", track)
	return FlightData{Source: track.Source, Modes: track.Modes, Callsign: track.Callsign,
		SSR: track.SSR, LastUpdate: track.LastUpdate}
}

func parse(file string) *template.Template {
	return template.Must(template.New("layout.html").ParseFS(files, "layout.html", file))
}
