package mailba

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMandrillSend(t *testing.T) {
	assert := assert.New(t)
	sender, err := NewMandrill(os.Getenv("MANDRILL_API_KEY"))
	assert.NoError(err)

	mail := NewMail()
	mail.SetFrom("witoo@plimble.com", "Tester")
	mail.SetSubject("hello world")
	mail.SetContent("<h1>hello</h1><p>*|data1|*</p><p>*|data2|*</p>")

	mail.AddGlobalVar("data1", "cat")
	mail.AddGlobalVar("data2", "dog")

	mail.AddMergeVar("xier.kokis@gmail.com", "data1", "art")
	mail.AddMergeVar("xier.kokis@gmail.com", "data2", "yo")

	mail.AddTo("xier.kokis@gmail.com", "Xier")
	mail.AddTo("witooh@icloud.com", "Jack")

	mail.AddAttachment("text/plain", "test.txt", base64.StdEncoding.EncodeToString([]byte("jack")))
	err = sender.Send(mail)
	assert.NoError(err)
}

func TestMandrillTemplateSend(t *testing.T) {
	assert := assert.New(t)
	sender, err := NewMandrill(os.Getenv("MANDRILL_API_KEY"))
	assert.NoError(err)

	mail := NewMail()
	mail.SetFrom("witoo@plimble.com", "Tester")
	mail.SetTemplate("test_offer")
	mail.SetSubject("hello world")
	mail.AddTo("xier.kokis@gmail.com", "Xier")
	mail.AddAttachment("text/plain", "test.txt", base64.StdEncoding.EncodeToString([]byte("jack")))
	err = sender.Send(mail)
	assert.NoError(err)
}
