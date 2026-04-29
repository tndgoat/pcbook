package user

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	Country     string `json:"country"`
}

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Contact   []Contact `json:"contact"`
}
