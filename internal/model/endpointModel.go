package model

type EndpointModel struct {
	ID      string      `json:"-" gorm:"primaryKey; autoIncrement:false"`
	Name    string      `json:"name"`
	Headers string      `json:"headers,omitempty" gorm:"optional"`
	URL     string      `json:"url"`
	Bodies  []BodyModel `json:"-" gorm:"foreignKey:EndpointID;constraint:OnDelete:CASCADE;"`
}
