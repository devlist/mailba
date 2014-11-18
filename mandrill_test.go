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

	mail := NewMail("witoo@plimble.com", "Tester")
	mail.SetSubject("hello world")
	mail.SetContent("<h1>hello</h1><p>world</p>")
	mail.AddTo("witooh@icloud.com", "Witoo")
	mail.AddAttachment("text/plain", "test.txt", base64.StdEncoding.EncodeToString([]byte("jack")))
	err = sender.Send(mail, nil)
	assert.NoError(err)
}

func TestMandrillTemplateSend(t *testing.T) {
	assert := assert.New(t)
	sender, err := NewMandrill("Ai4SC5Uv5BQ7gc6N2gvKPA")
	assert.NoError(err)

	mail := NewMail("witoo@plimble.com", "Tester")
	mail.SetTemplate("test_offer")
	mail.SetSubject("hello world")
	mail.AddTo("witooh@icloud.com", "Witoo")
	mail.AddAttachment("text/plain", "test.txt", base64.StdEncoding.EncodeToString([]byte("jack")))
	err = sender.Send(mail, nil)
	assert.NoError(err)
}
