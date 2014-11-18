package mailba

type recipient struct {
	email string
	name  string
}

type recipients map[string]*recipient

func newRecipientList() recipients {
	return make(recipients)
}

func (r recipients) add(email, name string) {
	r[email] = &recipient{email, name}
}
