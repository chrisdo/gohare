package storage

import "github.com/chrisdo/gohare/track"

type MsgType string

const (
	UPDATE MsgType = "UPDATE"
	REMOVE MsgType = "REMOVE"
	INSERT MsgType = "INSERT"
)

type ModelChangeMsg struct {
	Type MsgType
}

type Storage interface {
	StorageReader
	RemoveTrackById(modes string) bool
	InsertTrack(m *track.Model)
	CleanUp()
	SubscribeForChanges(c chan ModelChangeMsg)
}

type StorageReader interface {
	GetTrackById(modes string) (*track.Model, error)
	GetAllTracks() []track.Model
}
