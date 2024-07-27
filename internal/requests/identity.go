package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Signup struct {
	Username   string `form:"username" json:"username"`
	Password   string `form:"password" json:"password"`
	RePassword string `form:"repassword" json:"repassword"`
}

func (r *Signup) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}
