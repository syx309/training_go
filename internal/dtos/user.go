package dtos

type User struct {
	Id       string
	Name     string
	Email 	 string
	Password string
}

type UserData struct {
	Email string
}

type Login struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}

type Register struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
	Name     	string `json:"name"`
}
