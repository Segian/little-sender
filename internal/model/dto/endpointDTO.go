package dto

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type EndpointDto struct {
	Name    string         `json:"name"`
	Method  string         `json:"method"`
	Headers string         `json:"headers,omitempty"`
	URL     string         `json:"url"`
	Body    datatypes.JSON `json:"body,omitempty"`
}

type EndpointDtoUpdate struct {
	Name    string         `json:"name,omitempty"`
	Method  string         `json:"method,omitempty"`
	Headers string         `json:"headers,omitempty"`
	URL     string         `json:"url,omitempty"`
	Body    datatypes.JSON `json:"body,omitempty"`
}

type EndpointResponseDto struct {
	ID      uuid.UUID      `json:"id"`
	Name    string         `json:"name"`
	Method  string         `json:"method"`
	Headers string         `json:"headers,omitempty"`
	URL     string         `json:"url"`
	Body    datatypes.JSON `json:"body,omitempty"`
}
