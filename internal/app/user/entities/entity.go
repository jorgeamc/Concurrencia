package entity

type User struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Msg struct {
	Value string `json:"value" binding:"required"`
}
