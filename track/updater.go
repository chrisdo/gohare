package track

import (
	"github.com/chrisdo/gohare/opensky"
)

func (t *Model) UpdateByOpenSky(s *opensky.StateVector) {
	if len(s.Callsign) > 0 {
		t.Callsign = s.Callsign
	}
	tp, ok := t.LastTp()
	if !ok || (ok && s.TimePosition.After(tp.Timestamp)) {

		newTp := TrajectoryPoint{Timestamp: s.TimePosition, LocationWithAlt: LocationWithAlt{Location: Location{Lat: s.Latitude, Lon: s.Longitude}, Alt: int(s.AltitudeGeo)},
			Velocity: Velocity{Heading: s.TrueTrack, Speed: s.Velocity, VerticalRate: int32(s.VerticalRate)}}
		t.AddTrajectoryPoint(newTp)
	}
	if len(s.Squawk) > 0 {
		t.SSR = s.Squawk
	}
	t.LastUpdate = s.TimePosition

}
