package poker_time

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemPlayersStore struct {
	database *json.Encoder
	league   League
}

// NewFileSystemPlayersStore constructor for the server which stores the poker league
func NewFileSystemPlayersStore(file *os.File) (*FileSystemPlayersStore, error) {
	err := initialisePlayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("could not initialise player db file %s: %v", dbFileName, err)
	}
	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("could not load players store from file %s: %v", file.Name(), err)
	}

	return &FileSystemPlayersStore{
		database: json.NewEncoder(file),
		league:   league,
	}, nil
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("could not get info from file %s: %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

type League []Player

func (l League) Find(name string) *Player {
	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}
	return nil
}

func (f *FileSystemPlayersStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Score > f.league[j].Score
	})
	return f.league
}

func (f *FileSystemPlayersStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Score
	}

	return 0
}

func (f *FileSystemPlayersStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Score++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	f.database.Encode(&f.league)
}
