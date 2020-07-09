package Models

import "time"

type Post struct {
	ID   		string 		`json:"ID" bson:"_id,omitempty"`
	Title 		string 		`json:"title" bson:"title" binding:"required"`
	Content 	string 		`json:"content" bson:"content" binding:"required,min=500"`
	PermaLink 	string 		`json:"permaLink" bson:"permaLink" binding:"required"`
	Status	 	string 		`json:"status" bson:"status"`
	CreatedDate time.Time 	`json:"createdDate" bson:"createdDate"`
	CreatedBy 	string 		`json:"createdBy" bson:"createdBy"`
}