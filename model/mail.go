package model

//Sender details
type sender struct {
	Name  string
	Email string
}

//Receipient details
type recipient struct {
	Name  string
	Email string
}

//Mail contains mail data
type Mail struct {
	Sender    sender
	Recipient recipient
	Subject   string
	Body      string
	HTMLBody  string
}
