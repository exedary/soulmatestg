package uuid

import "github.com/google/uuid"

// GenerateID generates a unique ID that can be used as an identifier for an entity.
func Generate() string {
	return uuid.New().String()
}
