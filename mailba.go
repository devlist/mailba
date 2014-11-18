package mailba

type Config map[string]interface{}

type Interface interface {
	Send(mail *Mail, config *Config) error
}
