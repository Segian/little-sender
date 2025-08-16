package model

import (
	"github.com/google/uuid"
)

type EndpointModel struct {
	ID      uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;"`
	Name    string     `json:"name"`
	Method  string     `json:"method"`
	Headers string     `json:"headers,omitempty" gorm:"optional"`
	URL     string     `json:"url"`
	Body    *BodyModel `json:"body,omitempty" gorm:"foreignKey:EndpointID;constraint:OnDelete:CASCADE;"`
}
