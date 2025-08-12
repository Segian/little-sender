package dto

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type EndpointDto struct {
	Name    string         `json:"name"`
	Headers string         `json:"headers,omitempty"`
	URL     string         `json:"url"`
	Body    datatypes.JSON `json:"body,omitempty"`
}

type EndpointResponseDto struct {
	ID      uuid.UUID      `json:"id"`
	Name    string         `json:"name"`
	Headers string         `json:"headers,omitempty"`
	URL     string         `json:"url"`
	Body    datatypes.JSON `json:"body,omitempty"`
}
