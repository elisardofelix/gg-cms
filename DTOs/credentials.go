package DTOs


type Credentials struct {
	UserName string `form:"userName" json:"userName"`
	Password string `form:"password" json:"password"`
}
