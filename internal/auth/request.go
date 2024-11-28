package auth

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Reg struct {
	Username   string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	RePassword string `form:"repassword" json:"repassword" binding:"required,eqfield=Password"`
}

type Auth struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (r *Reg) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Password, validation.Required),
		validation.Field(&r.RePassword, validation.Required),
	)
}
