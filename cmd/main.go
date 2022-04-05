package main

import (
	"github.com/chrisdo/gohare/opensky"
	"github.com/chrisdo/gohare/storage"
	"github.com/chrisdo/gohare/track"
	"github.com/chrisdo/gohare/web"
	"github.com/rs/zerolog/log"

	"bufio"
	"os"
	"time"
)

func main() {
	storage := storage.NewMemoryStorage()

	oskyChan := make(chan *opensky.StateVectorResponse)

	oskyReader := opensky.NewReader(5)
	oskyReader.SetBoundingBox(45.8389, 5.9962, 47.8229, 10.5226)
	defer close(oskyChan)

	ticker := time.NewTicker(time.Second)
	go func() {
		<-ticker.C
		storage.CleanUp()
	}()

	go oskyReader.Connect(oskyChan)

	go func() {
		for {
			select {
			case r := <-oskyChan:
				processOskyVector(storage, r)
			//we can have more cases, e.g. from different sources
			default:
				//log.Trace().Msg("Waiting for data")
			}
		}
	}()

	server := web.NewServer(":8080", storage)
	server.StartServer()
	//TODO: this should run in a go routine, but then we need something that blocks the code afterwards
	bufio.NewReader(os.Stdin).ReadBytes('\n')
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
