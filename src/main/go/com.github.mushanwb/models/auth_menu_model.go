package models

type AuthMenu struct {
	BaseModel

	MenuName string `gorm:"type:varchar(255); not null" json:"menu_name"`
	Mark     string `gorm:"type:varchar(255); not null; unique" json:"mark"`
}
