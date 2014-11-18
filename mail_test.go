package mailba

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMail(t *testing.T) {
	assert := assert.New(t)
	mail := NewMail("Tester", "witoo@plimble.com")
	assert.Equal(recipient{"Tester", "witoo@plimble.com"}, mail.from)

	mail.SetSubject("subject")
	assert.Equal("subject", mail.subject)

	mail.SetTemplate("template")
	assert.Equal("template", mail.template)

	mail.SetContent("<html>")
	assert.Equal("<html>", mail.content)

	mail.SetFrom("John Doe", "test@test.com")
	assert.Equal(recipient{"test@test.com", "John Doe"}, mail.from)

	mail.to.add("to1@to.com", "To1")
	assert.Equal(&recipient{"to1@to.com", "To1"}, mail.to["to1@to.com"])
	assert.Len(mail.to, 1)

	mail.to.add("to2@to.com", "To2")
	assert.Equal(&recipient{"to2@to.com", "To2"}, mail.to["to2@to.com"])
	assert.Len(mail.to, 2)

	mail.cc.add("cc1@cc.com", "cc1")
	assert.Equal(&recipient{"cc1@cc.com", "cc1"}, mail.cc["cc1@cc.com"])
	assert.Len(mail.cc, 1)

	mail.cc.add("cc2@cc.com", "cc2")
	assert.Equal(&recipient{"cc2@cc.com", "cc2"}, mail.cc["cc2@cc.com"])
	assert.Len(mail.cc, 2)

	mail.bcc.add("bcc1@bcc.com", "bcc1")
	assert.Equal(&recipient{"bcc1@bcc.com", "bcc1"}, mail.bcc["bcc1@bcc.com"])
	assert.Len(mail.bcc, 1)

	mail.bcc.add("bcc2@bcc.com", "bcc2")
	assert.Equal(&recipient{"bcc2@bcc.com", "bcc2"}, mail.bcc["bcc2@bcc.com"])
	assert.Len(mail.bcc, 2)

	mail.AddAttachment("text/plain", "test.txt", "hello world")
	assert.Equal(&file{"text/plain", "test.txt", "hello world"}, mail.files[0])
	assert.Len(mail.files, 1)

	mail.AddAttachment("text/plain", "test2.txt", "hello world")
	assert.Equal(&file{"text/plain", "test2.txt", "hello world"}, mail.files[1])
	assert.Len(mail.files, 2)
}
