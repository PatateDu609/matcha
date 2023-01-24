package payloads

import (
	"fmt"
	"net/http"

	"github.com/PatateDu609/matcha/utils/database"
	"github.com/google/uuid"
)

type Room struct {
	ID    uuid.UUID `json:"uuid"`
	User1 uuid.UUID `json:"user1"`
	User2 uuid.UUID `json:"user2"`
}

func (room *Room) GetName() string {
	return "rooms"
}

func (room *Room) GetAlias() string {
	return "ro"
}

func (room *Room) GetColumns() []string {
	return database.GetColumns[Room](true)
}

func (room *Room) GetMandatoryColumns() []string {
	return database.GetColumns[Room](false)
}

func (room *Room) PrepareInsertion() ([]string, []any) {
	keys := room.GetColumns()
	return keys, database.PrepareValues(room)
}

func NewRoom(r *http.Request, user1, user2 uuid.UUID) (*Room, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("couldn't get uuid: %v", err)
	}

	room := &Room{
		ID:    id,
		User1: user1,
		User2: user2,
	}

	if err := database.Insert(r.Context(), room); err != nil {
		return nil, fmt.Errorf("couldn't insert room: %+v", err)
	}
	return room, nil
}
