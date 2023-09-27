package domain

type Person struct {
	UUID     string `json:"uuid"`
	FistName string `json:"fist_name"`
	Email    string `json:"email"`
	Ege      int    `json:"ege"`
}
