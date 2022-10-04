package structures

type UserStruct struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	EmailCode string `json:"email_code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type NoteStruct struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	ExpTime   string `json:"exp_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type JwtRequestModel struct {
	Token string `json:"token"`
}

type UserCheckRequest struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type EmailVer struct {
	Email     string `json:"email"`
	EmailCode string `json:"email_code"`
}
