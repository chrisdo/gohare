package storage

import (
	"fmt"
	"sync"
	"time"

	"github.com/chrisdo/gohare/track"
	"github.com/rs/zerolog/log"
)

type MemoryStorage struct {
	tracks map[string]*track.Model
	sync.RWMutex
	subscribers []chan ModelChangeMsg
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{tracks: make(map[string]*track.Model, 50)}
}

func (ms *MemoryStorage) CleanUp() {
	//should be locked
	now := time.Now()
	ms.Lock()
	defer ms.Unlock()
	for k, v := range ms.tracks {
		d := now.Sub(v.LastUpdate)
		if int64(d/time.Second) > 120 {
			delete(ms.tracks, k)
			log.Info().Str("modes", k).Str("reason", "timeout").Msg("Track removed")
		}
	}
}

func (ms *MemoryStorage) GetTrackById(modes string) (*track.Model, error) {
	//log.WithField("modes", modes).Info("Get Track By Id")
	ms.Lock()
	defer ms.Unlock()

	if _, ok := (ms.tracks[modes]); ok {
		return ms.tracks[modes], nil
	}
	log.Error().Str("modes", modes).Msg("Track not present")
	return nil, fmt.Errorf("Modesid %s not present in store", modes)
}

func (ms *MemoryStorage) RemoveTrackById(modes string) bool {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := (ms.tracks[modes]); !ok {
		return false
	}
	delete(ms.tracks, modes)
	return true
}

func (ms *MemoryStorage) InsertTrack(m *track.Model) {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := (ms.tracks[m.Modes]); ok {
		log.Warn().Str("modes", m.Modes).Msg("Track with modes ID already present, will override")

	}
	ms.tracks[m.Modes] = m
}

func (ms *MemoryStorage) GetAllTracks() []track.Model {
	ms.Lock()
	defer ms.Unlock()

	tracks := make([]track.Model, 0)
	for _, v := range ms.tracks {
		if v.Modes != "" {
			tracks = append(tracks, *v)
		}
	}
	return tracks
}

func (ms *MemoryStorage) SubscribeForChanges(c chan ModelChangeMsg) {
	ms.subscribers = append(ms.subscribers, c)
}
