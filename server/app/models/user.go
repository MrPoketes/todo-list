package models

type User struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Items    []Items `json:"items"`
}

type Items struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}
