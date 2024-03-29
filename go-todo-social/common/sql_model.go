package common

import "time"

type SQLModel struct {
	Id int `json:"id" gorm:"column:id;"`
	CreatedAt   *time.Time  `json:"created_at" gorm:"created_at;"`
	UpadatedAt  *time.Time  `json:"updated_at" gorm:"updated_at;"`
}
