package mailba

//go:generate mockgen -destination=mock.go --self_package=github.com/plimble/mailba -package=mailba github.com/plimble/mailba Sender

type Config map[string]interface{}

type Sender interface {
	Send(mail *Mail, config *Config) error
}
