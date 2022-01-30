package models

import "github.com/arathi/wealth-manager/internal/forms"

type Data struct {
	StringMap    map[string]string
	DataMap      map[string]interface{}
	LoggedInUser User
	Form         *forms.Form
}
