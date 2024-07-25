package requests

type Signup struct {
	Username   string `form:"username" json:"username"`
	Password   string `form:"password" json:"password"`
	RePassword string `form:"repassword" json:"repassword"`
}
