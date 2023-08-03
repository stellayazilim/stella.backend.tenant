package Types

type UserLoginWithCredentialRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
