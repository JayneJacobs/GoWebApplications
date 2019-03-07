package viewmodel

// Login defines user accout
type Login struct {
	Title    string
	Active   string
	Email    string
	Password string
}

//NewLogin defines login Function
func NewLogin() Login {
	result := Login{
		Active: "home",
		Title:  "Band Login",
	}
	return result
}
