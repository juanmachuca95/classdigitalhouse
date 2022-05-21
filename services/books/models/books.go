package models

type Book struct {
	Id         int    `json:"id"`
	Book       string `json:"book"`
	Author     string `json:"author"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
