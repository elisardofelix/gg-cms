package Models

import "time"

type User struct {
	ID   		string 		`json:"ID" bson:"_id,omitempty"`
	Username	string 		`json:"username" bson:"username " binding:"required"`
	Password 	string 		`json:"password" bson:"password" binding:"required"`
	Email 		string 		`json:"email" bson:"email" binding:"required,email" `
	CreatedDate time.Time 	`json:"createdDate" bson:"createdDate"`
	CreatedBy 	string 		`json:"createdBy" bson:"createdBy"`
}