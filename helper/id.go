package helper

import "github.com/google/uuid"

func GeneratedUUID() uuid.UUID {
	newUUID := uuid.New()
	return newUUID
}
