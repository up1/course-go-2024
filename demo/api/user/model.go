package user

import "time"

type UserResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CreatedDate time.Time `json:"created_date"`
}
