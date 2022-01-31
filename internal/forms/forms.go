package forms

import (
	"net/url"

	"github.com/arathi/wealth-manager/internal/config"
	"github.com/go-playground/validator/v10"
)

var app *config.AppConfig

// errors hold errors for fields
type errors map[string][]string

// Form is a custom struct to hold form data and errors (if any)
type Form struct {
	url.Values
	Errors errors
}

func Init(a *config.AppConfig) {
	app = a
	app.Validate = validator.New()
}

// Init initializes new Form
func New(data url.Values) *Form {
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

// Valid returns true if form contain no errors, false otherwise
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// ValidateField runs validation on field and adds errors into Form
func (f *Form) ValidateField(validation, fieldName, fieldValue, errMessage string) {
	err := app.Validate.Var(fieldValue, validation)
	if err != nil {
		if errMessage == "" {
			errMessage = err.Error()
		}
		f.Errors.Add(fieldName, errMessage)
	}

}
