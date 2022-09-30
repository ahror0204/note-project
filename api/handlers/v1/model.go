package v1

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