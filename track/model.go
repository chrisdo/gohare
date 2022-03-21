package track

import (
	"fmt"
	"time"
)

type SpeedIndicator int

const (
	GS SpeedIndicator = iota + 1
	TAS
	IAS
)

const (
	ADSB    string = "ADSB"
	ASTERIX string = "ASTERIX"
	XPLANE  string = "XPlane"
	OPENSKY string = "OpenSKy"
)

//TODO abstract the model a bit further so its not so ADS-B related but more general, e.g. CPR, could do a ads-b track containing a Model and cpr
//and a e.g. Xplane Track, which would not contain the cpr
type Model struct {
	Source             string
	Modes              string
	EmitterCategory    string
	SurveillanceStatus byte
	Callsign           string
	SSR                string
	Mops               byte
	LastUpdate         time.Time
	History            []TrajectoryPoint
	vel                Velocity
	lastCprData        CprData //TODO this should go to adsb beause its only part of adsb decoding
}

type Velocity struct {
	SpeedType    SpeedIndicator
	Speed        float64
	Heading      float64
	VerticalRate int32
	Timestamp    time.Time
}

type TrajectoryPoint struct {
	LocationWithAlt
	Velocity
	Timestamp time.Time
}

type Location struct {
	Lat float64
	Lon float64
}

type LocationWithAlt struct {
	Location
	Alt int
}

type CprData struct {
	Odd         bool
	Lat         int32
	Lon         int32
	LastCprTime time.Time
}

func NewModel(modes, source string) *Model {
	return &Model{Modes: modes, Source: source, History: make([]TrajectoryPoint, 0, 100)}
}

func (m *Model) Position() (loc LocationWithAlt, ok bool) {
	if len(m.History) == 0 {
		return LocationWithAlt{}, false
	}
	return m.History[len(m.History)-1].LocationWithAlt, true
}

func (m *Model) LastTp() (tp TrajectoryPoint, ok bool) {
	if len(m.History) == 0 {
		return TrajectoryPoint{}, false
	}
	return m.History[len(m.History)-1], true
}

func (m *Model) UpdateVelocity(v Velocity) {
	m.vel = v
}

func (m *Model) Velocity() Velocity {
	return m.vel
}

func (m *Model) AddTrajectoryPoint(t TrajectoryPoint) {
	m.History = append(m.History, t)
}

func (m *Model) AddLocation(t time.Time, l LocationWithAlt) {
	m.History = append(m.History, TrajectoryPoint{Timestamp: t, LocationWithAlt: l, Velocity: Velocity{SpeedType: m.vel.SpeedType, Speed: m.vel.Speed, Heading: m.vel.Heading, VerticalRate: m.vel.VerticalRate}})
}

func (m *Model) UpdateCpr(odd bool, lat, lon int32, t time.Time) {
	m.lastCprData = CprData{odd, lat, lon, t}

}

func (m *Model) LastCpr() CprData {
	return m.lastCprData
}

func (m *Model) String() string {
	pos, _ := m.Position()
	return fmt.Sprintf("\n%s %s %f %f\n", m.LastUpdate, m.Modes, pos.Lat, pos.Lon)
}
