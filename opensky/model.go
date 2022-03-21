package opensky

import (
	"encoding/json"
	"time"
)

//StateVector of an OpenSky Track, describes the current state of the track
type StateVector struct {
	Icao24         string
	Callsign       string
	OriginCountry  string
	TimePosition   time.Time
	LastContact    time.Time
	Longitude      float64
	Latitude       float64
	AltitudeBaro   float64
	OnGround       bool
	Velocity       float64
	TrueTrack      float64
	VerticalRate   float64
	Sensors        []int
	AltitudeGeo    float64
	Squawk         string
	Spi            bool
	PositionSource int
}

//StateVectorResponse contains a time and a set of StateVectors
type StateVectorResponse struct {
	Time   int           `json:"time"`
	States []StateVector `json:"states"`
}

const metersToFeet float64 = 3.28084
const mPerSToFtPerMin float64 = 196.85
const mPerSToKnots float64 = 1.94384

//UnmarshalResponse to unmarshal the json data response into the given StateVectorResponse
func (s *StateVectorResponse) UnmarshalResponse(data []byte) error {
	return json.Unmarshal(data, &s)
}

//UnmarshalJSON unmarshal the raw json data into a given StateVector
func (s *StateVector) UnmarshalJSON(data []byte) error {

	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	s.Icao24, _ = v[0].(string)
	if v[1] != nil {
		s.Callsign, _ = v[1].(string)
	}
	s.OriginCountry = v[2].(string)
	if v[3] != nil {
		s.TimePosition = time.Unix(int64(v[3].(float64)), 0)
	}
	s.LastContact = time.Unix(int64(v[4].(float64)), 0)
	if v[5] != nil {

		s.Longitude = v[5].(float64)
	}
	if v[6] != nil {

		s.Latitude = v[6].(float64)
	}
	if v[7] != nil {

		s.AltitudeBaro = v[7].(float64) * metersToFeet
	}
	s.OnGround = v[8].(bool)
	if v[9] != nil {

		s.Velocity = v[9].(float64) * mPerSToKnots
	}
	if v[10] != nil {

		s.TrueTrack = v[10].(float64)
	}
	if v[11] != nil {

		s.VerticalRate = v[11].(float64) * mPerSToFtPerMin
	}
	//v[12] is sensoes, skip
	if v[13] != nil {

		s.AltitudeGeo = v[13].(float64) * metersToFeet
	}
	if v[14] != nil {
		s.Squawk = v[14].(string)
	}
	s.Spi = v[15].(bool)
	s.PositionSource = int(v[16].(float64))
	return nil
}
