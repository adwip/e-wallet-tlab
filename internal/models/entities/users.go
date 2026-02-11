package entities

import "time"

type Users struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement;index"`
	SecureId  string    `gorm:"column:secure_id;not null;uniqueIndex;type:varchar(255)"`
	Name      string    `gorm:"column:name;not null;type:varchar(255)"`
	Email     string    `gorm:"column:email;not null;uniqueIndex;type:varchar(255)"`
	Password  string    `gorm:"column:password;not null;type:varchar(255)"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;index"`
}

func (Users) TableName() string {
	return "users"
}
