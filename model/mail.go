package model

//Sender details
type sender struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

//Receipient details
type recipient struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

//Mail contains mail data
type Mail struct {
	Sender    sender    `json:"sender" validate:"required,email"`
	Recipient recipient `json:"recipient" validate:"required,email"`
	Subject   string    `json:"subject" validate:"required"`
	Body      string    `json:"body" validate:"required"`
	HTMLBody  string    `json:"htmlbody" validate:"required"`
}
