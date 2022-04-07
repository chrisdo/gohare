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
	pb "github.com/chrisdo/gohare/trackstoreservice"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
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
	client        pb.TrackStoreServiceClient
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

func NewServer(addr string, client pb.TrackStoreServiceClient) *Server {
	fileServer := http.FileServer(http.Dir("./Static"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	mux.Handle("/resources/", http.StripPrefix("/resources", fileServer))
	mux.HandleFunc("/flights", generalFlightsHandler(client))
	return &Server{http.Server{Addr: addr, Handler: mux}, nil, client}

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

func generalFlightsHandler(client pb.TrackStoreServiceClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fid := r.URL.Query().Get("fid")
		if len(fid) == 0 {
			log.Info().Msg("Handling flight list")
			err := flightsHandler(w, getFlightList(client))
			if err != nil {
				log.Err(err)
			}
		} else {
			log.Info().Str("ModesId", fid).Msg("Handling single flight")
			err := singleFlightHandler(w, getSingleFlight(client, fid))
			if err != nil {
				log.Err(err)
			}
		}
	}
}

func getFlightList(client pb.TrackStoreServiceClient) FlightList {
	tracks, err := client.GetFlightList(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Err(err)
		return FlightList{Flights: make([]FlightDataShort, 0)}
	}
	log.Debug().Msgf("Got %d tracks\n", len(tracks.Flights))
	result := make([]FlightDataShort, len(tracks.Flights))
	for i, track := range tracks.Flights {
		result[i] = FlightDataShort{Modes: track.ModesId, Callsign: track.Callsign, SSR: track.SSR}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[j].Modes < result[i].Modes
	})
	log.Debug().Msgf("Tracks: %v\n", result)
	return FlightList{nil, result}
}

func getSingleFlight(client pb.TrackStoreServiceClient, fid string) FlightData {
	track, err := client.GetFlight(context.Background(), &pb.FlightRequest{ModesId: fid})
	if err != nil {
		log.Err(err)
		return FlightData{Error: err}
	}
	log.Debug().Str("Modesid", fid).Msgf("%v", track)
	return FlightData{Source: "OpenSky", Modes: track.ModesId, Callsign: track.Callsign,
		SSR: track.SSR, LastUpdate: time.UnixMilli(track.LastUpdate)}
}

func parse(file string) *template.Template {
	return template.Must(template.New("layout.html").ParseFS(files, "layout.html", file))
}
