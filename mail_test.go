package mailba

import (
	"github.com/mattbaird/gochimp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMail(t *testing.T) {
	assert := assert.New(t)
	mail := NewMail("Tester", "witoo@plimble.com")
	assert.Equal(Recipient{"Tester", "witoo@plimble.com"}, mail.From)

	mail.SetSubject("subject")
	assert.Equal("subject", mail.Subject)

	mail.SetTemplate("template")
	assert.Equal("template", mail.Template)

	mail.SetContent("<html>")
	assert.Equal("<html>", mail.Content)

	mail.SetFrom("John Doe", "test@test.com")
	assert.Equal(Recipient{"test@test.com", "John Doe"}, mail.From)

	mail.AddTo("to1@to.com", "To1")
	assert.Equal(&Recipient{"to1@to.com", "To1"}, mail.To[0])
	assert.Len(mail.To, 1)

	mail.AddTo("to2@to.com", "To2")
	assert.Equal(&Recipient{"to2@to.com", "To2"}, mail.To[1])
	assert.Len(mail.To, 2)

	mail.AddCC("cc1@cc.com", "cc1")
	assert.Equal(&Recipient{"cc1@cc.com", "cc1"}, mail.CC[0])
	assert.Len(mail.CC, 1)

	mail.AddCC("cc2@cc.com", "cc2")
	assert.Equal(&Recipient{"cc2@cc.com", "cc2"}, mail.CC[1])
	assert.Len(mail.CC, 2)

	mail.AddBCC("bcc1@bcc.com", "bcc1")
	assert.Equal(&Recipient{"bcc1@bcc.com", "bcc1"}, mail.BCC[0])
	assert.Len(mail.BCC, 1)

	mail.AddBCC("bcc2@bcc.com", "bcc2")
	assert.Equal(&Recipient{"bcc2@bcc.com", "bcc2"}, mail.BCC[1])
	assert.Len(mail.BCC, 2)

	mail.AddAttachment("text/plain", "test.txt", "hello world")
	assert.Equal(&File{"text/plain", "test.txt", "hello world"}, mail.Files[0])
	assert.Len(mail.Files, 1)

	mail.AddAttachment("text/plain", "test2.txt", "hello world")
	assert.Equal(&File{"text/plain", "test2.txt", "hello world"}, mail.Files[1])
	assert.Len(mail.Files, 2)

	mail.AddGlobalVar("firstname", "art")
	mail.AddGlobalVar("catname", "silverbag")
	assert.Equal(Vars{
		gochimp.Var{"firstname", "art"},
		gochimp.Var{"catname", "silverbag"},
	}, mail.Globalvars)
	assert.Len(mail.Globalvars, 2)

	mail.AddMergeVar("witoohxx@gmail.com", "firstname", "art")
	mail.AddMergeVar("witoohxx@gmail.com", "catname", "silverbag")
	mail.AddMergeVar("momoxx@gmail.com", "firstname", "momo")
	mail.AddMergeVar("momoxx@gmail.com", "catname", "goldbag")
	assert.Len(mail.Mergevars, 2)
	assert.Equal(map[string]Vars{
		"witoohxx@gmail.com": []gochimp.Var{
			gochimp.Var{"firstname", "art"},
			gochimp.Var{"catname", "silverbag"},
		},

		"momoxx@gmail.com": Vars{
			gochimp.Var{"firstname", "momo"},
			gochimp.Var{"catname", "goldbag"},
		},
	}, mail.Mergevars)

}
