package profileEntity

import (
	"fmt"
	"strings"
)

type Profile struct {
	id         *ProfileID
	firstName  string
	lastName   string
	patronymic string
}

func NewEmptyProfile() *Profile {
	return NewBuilder().Build()
}

func (p Profile) ID() *ProfileID {
	return p.id
}

func (p Profile) FullName() string {
	fullName := fmt.Sprintf("%s %s", p.firstName, p.lastName)
	return strings.TrimSpace(fullName)
}

func (p Profile) LastName() string {
	return p.lastName
}

func (p Profile) FirstName() string {
	return p.firstName
}

func (p Profile) Patronymic() string {
	return p.patronymic
}

func (p *Profile) SetFirstName(firstName string) error {
	if firstName == "" {
		return ErrFirstNameIsRequired
	}

	p.firstName = firstName
	return nil
}

func (p *Profile) SetLastName(lastName string) error {
	if lastName == "" {
		return ErrLastNameIsRequired
	}

	p.lastName = lastName
	return nil
}

func (p *Profile) SetPatronymic(patronymic string) {
	p.patronymic = patronymic
}
