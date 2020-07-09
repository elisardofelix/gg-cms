package DTOs

type Registration struct {
	UserName	string 		`json:"username" binding:"required"`
	Password 	string 		`json:"password" binding:"required"`
	RePasword   string      `json:"repassword" binding:"required"`
	Email 		string 		`json:"email" binding:"required,email"`
	IsAdmin     bool 		`json:"isAdmin"`
	Status		string 		`json:"status" binding:"required"`
	CreatedBy	string 		`json:"createdBy"`
}