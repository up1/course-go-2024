package user

import "time"

type UserResponse struct {
	ID          int       
	Name        string    
	CreatedDate time.Time 
}
