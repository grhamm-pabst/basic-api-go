package entity

import "github.com/google/uuid"

type Todo struct {
	Id          uuid.UUID
	Description string
	Finished    bool
}
