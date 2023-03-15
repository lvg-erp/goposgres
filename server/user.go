package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"usernmae"`
	Password string `json:"password"`
}
