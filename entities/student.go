package entities

import "time"

type Students struct {
	ID         uint      `json:"id" gorm:"primary_key;autoIncrement"`
	FirstName  string    `json:"first_name" gorm:"not null"`
	LastName   string    `json:"last_name" gorm:"not null"`
	Phone      string    `json:"phone" gorm:"not null;unique"`
	Email      string    `json:"email" gorm:"not null;unique"`
	Birthday   time.Time `json:"birthday" gorm:"default:CURRENT_TIMESTAMP"`
	Gender     int8      `json:"gender" gorm:"not null;default:0"`
	University string    `json:"university"`
	Address    string    `json:"address"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
