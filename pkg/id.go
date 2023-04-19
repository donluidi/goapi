package pkg

import "github.com/google/uuid"

type Id = uuid.UUID

func NewId() Id {
	return Id(uuid.New())
}

func ParseId(s string) (Id, error) {
	_id, err := uuid.Parse(s)
	return Id(_id), err
}
