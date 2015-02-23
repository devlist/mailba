package mailba

//go:generate mockgen -destination=mock_mailba/mock_sender.go github.com/plimble/mailba Sender

type Config map[string]interface{}

type Sender interface {
	Send(mail *Mail, config *Config) error
}
