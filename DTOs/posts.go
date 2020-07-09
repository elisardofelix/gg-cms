package DTOs

import "gg-cms/Models"

type Posts struct {
	Data	[]Models.Post	`json:"data"`
	Total 	int64			`json:"total"`
}
