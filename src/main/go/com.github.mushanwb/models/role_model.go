package models

type Role struct {
	BaseModel

	RoleName string `gorm:"type:varchar(64); not null;" json:"role_name"`
}
