package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(database io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(database).Decode(&league)
	if err != nil {
		err = fmt.Errorf("league not readable as JSON: %v", err)
	}
	return league, err
}
