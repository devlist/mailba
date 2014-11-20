package mailba

import (
	"github.com/mattbaird/gochimp"
)

type Mail struct {
	from       recipient
	to         recipients
	cc         recipients
	bcc        recipients
	files      []*file
	content    string
	subject    string
	template   string
	mergevars  map[string][]gochimp.Var
	globalvars []gochimp.Var
}

func NewMail(defaultFromEmail, defaultFromName string) *Mail {
	return &Mail{
		to:        newRecipientList(),
		cc:        newRecipientList(),
		bcc:       newRecipientList(),
		files:     []*file{},
		from:      recipient{defaultFromEmail, defaultFromName},
		mergevars: make(map[string][]gochimp.Var),
	}
}

func (m *Mail) SetContent(html string) *Mail {
	m.content = html
	return m
}

func (m *Mail) SetSubject(subject string) *Mail {
	m.subject = subject
	return m
}

func (m *Mail) SetFrom(name, email string) *Mail {
	m.from = recipient{email, name}
	return m
}

func (m *Mail) AddTo(email, name string) *Mail {
	m.to.add(email, name)
	return m
}

func (m *Mail) AddCC(email, name string) *Mail {
	m.cc.add(email, name)
	return m
}

func (m *Mail) AddBCC(email, name string) *Mail {
	m.bcc.add(email, name)
	return m
}

func (m *Mail) AddAttachment(mime, name, content string) *Mail {
	m.files = append(m.files, &file{mime, name, content})
	return m
}

func (m *Mail) SetTemplate(name string) *Mail {
	m.template = name
	return m
}

func (m *Mail) AddMergeVar(email, key string, val string) *Mail {
	m.mergevars[email] = append(m.mergevars[email], gochimp.Var{key, val})
	return m
}

func (m *Mail) AddGlobalVar(key string, val string) *Mail {
	m.globalvars = append(m.globalvars, gochimp.Var{key, val})
	return m
}
