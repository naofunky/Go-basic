package pointer

type User struct {
	Name string `json:"name"`
}

func NewUser() User {
	aUser := User{Name: "John"}
	return aUser
}
