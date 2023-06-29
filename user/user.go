package user

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
