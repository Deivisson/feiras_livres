package domain

type Validation struct {
	Errors map[string][]string
}

func (v *Validation) AddErrorField(fieldName string, errors []string) {
	if len(errors) == 0 {
		return
	}
	if v.Errors == nil {
		v.Errors = map[string][]string{}
	}
	v.Errors[fieldName] = errors
}

func (v *Validation) HasError() bool {
	return len(v.Errors) > 0
}
