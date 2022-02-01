package models

import "github.com/arathi/wealth-manager/internal/forms"

type Data struct {
	StringMap     map[string]string
	DataMap       map[string]interface{}
	LoggedInUser  User
	Form          *forms.Form
	CSRFToken     string
	Authenticated bool
	ErrorMsg      string
	SuccessMsg    string
}
