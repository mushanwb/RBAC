package models

type AuthMenu struct {
	BaseModel

	Name string `gorm:"type:varchar(255); not null" json:"name"`
	Mark string `gorm:"type:varchar(255); not null; unique" json:"mark"`
}
