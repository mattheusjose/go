package functions

import "github.com/google/uuid"

func RandomUUID() uuid.UUID {

	return uuid.New()
}
