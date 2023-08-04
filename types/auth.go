package Types

type UserLoginWithCredentialRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserRequestBody struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func (r *RegisterUserRequestBody) ConvertToUser() *User {

	return &User{
		Name:        r.Name,
		Email:       r.Email,
		PhoneNumber: r.PhoneNumber,
		Password:    []byte(r.Password),
	}
}
