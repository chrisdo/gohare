package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/chrisdo/gohare/opensky"
	"github.com/chrisdo/gohare/storage"
	"github.com/chrisdo/gohare/track"
	pb "github.com/chrisdo/gohare/trackstoreservice"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"
)

var store storage.Storage

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 50051, "The server port")
)

type trackStoreServer struct {
	store storage.StorageReader
	pb.UnimplementedTrackStoreServiceServer
}

func (s *trackStoreServer) GetFlightList(context.Context, *emptypb.Empty) (*pb.FlightList, error) {
	tracks := s.store.GetAllTracks()
	converted := make([]*pb.Flight, len(tracks))
	for i, t := range tracks {
		converted[i] = convertTrack(&t)
	}
	return &pb.FlightList{Flights: converted}, nil
}

func (s *trackStoreServer) GetFlight(ctx context.Context, r *pb.FlightRequest) (*pb.Flight, error) {
	t, err := s.store.GetTrackById(r.ModesId)
	if err != nil {
		return nil, err
	}
	return convertTrack(t), nil
}

func newTrackStoreServer() *trackStoreServer {
	return &trackStoreServer{store: store}
}

func convertTrack(t *track.Model) *pb.Flight {
	return &pb.Flight{ModesId: t.Modes, Callsign: t.Callsign, SSR: t.SSR, LastUpdate: t.LastUpdate.Unix()}
}

func main() {
	flag.Parse()
	store = storage.NewMemoryStorage()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatal().Msgf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterTrackStoreServiceServer(grpcServer, newTrackStoreServer())

	oskyChan := make(chan *opensky.StateVectorResponse)
	oskyReader := opensky.NewReader(5)
	oskyReader.SetBoundingBox(45.8389, 5.9962, 47.8229, 10.5226)
	defer close(oskyChan)

	ticker := time.NewTicker(time.Second)
	go func() {
		<-ticker.C
		store.CleanUp()
	}()

	go oskyReader.Connect(oskyChan)
	go func() {
		for {
			select {
			case r := <-oskyChan:
				processOskyVector(store, r)
			//we can have more cases, e.g. from different sources
			default:
				//log.Trace().Msg("Waiting for data")
			}
		}
	}()
	grpcServer.Serve(lis)

	//this is actually only really useful when having multiple readers, but lets keep it for learning
}

func processOskyVector(storage storage.Storage, r *opensky.StateVectorResponse) {
	log.Info().Msg(time.Unix(int64(r.Time), 0).String())
	for _, s := range r.States {
		m, err := storage.GetTrackById(s.Icao24)
		if err != nil {
			log.Error().Err(err)

			m = track.NewModel(s.Icao24, track.OPENSKY)
			storage.InsertTrack(m)
		}
		m.UpdateByOpenSky(&s)

	}
}
