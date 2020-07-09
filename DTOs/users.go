package DTOs

type Users struct {
	Data	[]User 	`json:"data"`
	Total	int64 	`json:"total"`
}