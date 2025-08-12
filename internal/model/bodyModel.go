package model

import "gorm.io/datatypes"

type BodyModel struct {
	EndpointID string         `json:"-" gorm:"primaryKey;not null"` //clave primaria y foranea
	Data       datatypes.JSON `json:"body" gorm:"type:json"`
}
