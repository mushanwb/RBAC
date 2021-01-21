package models

type User struct {
	BaseModel

	UserName string `gorm:"type:varchar(64); not null;" json:"user_name"`
	Phone    string `gorm:"type:varchar(11); not null; unique" json:"phone"`
	Password string `gorm:"type:varchar(255); not null;" json:"-"`
	RoleId   int8   `gorm:"type:tinyint(1); not null;" json:"role_id"`
}
