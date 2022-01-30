package forms

import "net/url"

// errors hold errors for fields
type errors map[string][]string

// Form is a custom struct to hold form data and errors (if any)
type Form struct {
	url.Values
	Errors errors
}

// Init initializes new Form
func Init(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Add adds error message for a field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get retrieves error message for a field if exists
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
