package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type EndpointModel struct {
	ID      uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;"`
	Name    string         `json:"name"`
	Method  string         `json:"method"`
	Headers datatypes.JSON `json:"headers,omitempty" gorm:"type:json;optional"`
	URL     string         `json:"url"`
	Body    *BodyModel     `json:"body,omitempty" gorm:"foreignKey:EndpointID;constraint:OnDelete:CASCADE;"`
}
