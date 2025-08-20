package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type HistoryModel struct {
	ID            uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;not null"`
	EndpointID    uuid.UUID      `json:"endpoint_id" gorm:"type:uuid;not null"`
	EndpointModel EndpointModel  `json:"-" gorm:"foreignKey:EndpointID;references:ID;constraint:OnDelete:CASCADE"`
	StatusCode    int            `json:"status_code" gorm:"not null"`
	Timestamp     time.Time      `json:"timestamp" gorm:"not null"`
	Response      datatypes.JSON `json:"response" gorm:"type:json;not null"`
	ResponseTime  int64          `json:"response_time" gorm:"not null"`
}
