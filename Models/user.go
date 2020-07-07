package Models

import "time"

type User struct {
	ID   		string 		`json:"ID" bson:"_id,omitempty"`
	UserName	string 		`json:"userName" bson:"userName" binding:"required"`
	Password 	string 		`json:"password" bson:"password"`
	Email 		string 		`json:"email" bson:"email" binding:"required,email" `
	Status	 	string 		`json:"status" bson:"status" binding:"required"`
	IsAdmin		bool		`json:"isAdmin" bson:"isAdmin" binding:"required"`
	CreatedDate time.Time 	`json:"createdDate" bson:"createdDate"`
	CreatedBy 	string 		`json:"createdBy" bson:"createdBy"`
}