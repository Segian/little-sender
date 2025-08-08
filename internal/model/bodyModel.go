package model

type BodyModel struct {
	ID         string `json:"id"`
	EndpointID string `json:"endpoint_id" gorm:"index;not null"`
	Body       string `json:"body"`
}
