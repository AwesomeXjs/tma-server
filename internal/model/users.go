package model

type User struct {
	ID        int    `json:"id" example:"51"`
	Username  string `json:"username" example:"username"`
	FirstName string `json:"firstName" example:"firstName"`
	LastName  string `json:"lastName" example:"lastName"`
	IsPremium int    `json:"isPremium" example:"0"`
}
