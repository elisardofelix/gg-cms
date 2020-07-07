package DTOs

import "time"

type User struct {
	ID          string    `json:"ID" bson:"_id,omitempty"`
	UserName    string    `json:"username" bson:"username " binding:"required"`
	Email       string    `json:"email" bson:"email" binding:"required,email"`
	Status      string    `json:"status" bson:"status" binding:"required"`
	IsAdmin		bool	  `json:"isAdmin" bson:"isAdmin" binding:"required"`
	CreatedDate time.Time `json:"createdDate" bson:"createdDate"`
	CreatedBy   string    `json:"createdBy" bson:"createdBy"`
}