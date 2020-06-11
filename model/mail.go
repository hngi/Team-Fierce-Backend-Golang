package model

//Sender details
type sender struct {
	name  string
	email string
}

//Receipient details
type recipient struct {
	name  string
	email string
}

//Mail contains mail data
type Mail struct {
	sender    sender
	recipient recipient
	subject   string
	body      string
	htmlBody  string
}
