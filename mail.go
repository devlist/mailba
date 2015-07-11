package mailba

import (
	"encoding/json"
	"github.com/mattbaird/gochimp"
	"github.com/plimble/utils/errors2"
	"gopkg.in/mgo.v2/bson"
)

//msgp:ignore File Files

type Recipient struct {
	Email string `json:"e" msg:"e"`
	Name  string `json:"n" msg:"n"`
}

type Recipients []*Recipient

type File struct {
	Mime    string
	Name    string
	Content string
}

type Files []*File

type Vars []gochimp.Var

type Mail struct {
	From       Recipient       `json:"f" msg:"f"`
	To         Recipients      `json:"to" msg:"to"`
	CC         Recipients      `json:"cc" msg:"cc"`
	BCC        Recipients      `json:"bcc" msg:"bcc"`
	Files      Files           `json:"-" msg:"-"`
	Content    string          `json:"c" msg:"c"`
	Subject    string          `json:"s" msg:"s"`
	Template   string          `json:"t" msg:"t"`
	Mergevars  map[string]Vars `json:"mv" msg:"mv"`
	Globalvars Vars            `json:"gv" msg:"gv"`
}

func NewMail(defaultFromEmail, defaultFromName string) *Mail {
	return &Mail{
		To:        Recipients{},
		CC:        Recipients{},
		BCC:       Recipients{},
		Files:     Files{},
		From:      Recipient{defaultFromEmail, defaultFromName},
		Mergevars: make(map[string]Vars),
	}
}

func (m *Mail) SetContent(html string) *Mail {
	m.Content = html
	return m
}

func (m *Mail) SetSubject(subject string) *Mail {
	m.Subject = subject
	return m
}

func (m *Mail) SetFrom(name, email string) *Mail {
	m.From = Recipient{email, name}
	return m
}

func (m *Mail) AddTo(email, name string) *Mail {
	m.To = append(m.To, &Recipient{email, name})
	return m
}

func (m *Mail) AddCC(email, name string) *Mail {
	m.CC = append(m.CC, &Recipient{email, name})
	return m
}

func (m *Mail) AddBCC(email, name string) *Mail {
	m.BCC = append(m.BCC, &Recipient{email, name})
	return m
}

func (m *Mail) AddAttachment(mime, name, content string) *Mail {
	m.Files = append(m.Files, &File{mime, name, content})
	return m
}

func (m *Mail) SetTemplate(name string) *Mail {
	m.Template = name
	return m
}

func (m *Mail) AddMergeVar(email, key string, val string) *Mail {
	m.Mergevars[email] = append(m.Mergevars[email], gochimp.Var{key, val})
	return m
}

func (m *Mail) AddGlobalVar(key string, val string) *Mail {
	m.Globalvars = append(m.Globalvars, gochimp.Var{key, val})
	return m
}

func (m *Mail) JSON() ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, errors2.NewInternal(err.Error())
	}

	return b, nil
}

func (m *Mail) BSON() ([]byte, error) {
	b, err := bson.Marshal(m)
	if err != nil {
		return nil, errors2.NewInternal(err.Error())
	}

	return b, nil
}
