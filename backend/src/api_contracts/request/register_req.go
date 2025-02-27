package request

type RegisterReq struct {
	Password  string
	Email     string
	Username  string
	FirstName string
	LastName  string
	Country   string
}