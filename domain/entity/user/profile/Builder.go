package profileEntity

type Builder struct {
	id         *ProfileID
	firstName  string
	lastName   string
	patronymic string
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) ID(id *ProfileID) *Builder {
	b.id = id
	return b
}

func (b *Builder) FirstName(firstName string) *Builder {
	b.firstName = firstName
	return b
}

func (b *Builder) LastName(lastName string) *Builder {
	b.lastName = lastName
	return b
}

func (b *Builder) Patronymic(patronymic string) *Builder {
	b.patronymic = patronymic
	return b
}

func (b *Builder) Build() *Profile {
	b.fillDefaultFields()

	return b.createFromBuilder()
}

func (b *Builder) fillDefaultFields() {
	if b.id == nil {
		b.id = NewProfileID()
	}
}

func (b *Builder) createFromBuilder() *Profile {
	return &Profile{
		id:         b.id,
		firstName:  b.firstName,
		lastName:   b.lastName,
		patronymic: b.patronymic,
	}
}
