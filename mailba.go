package mailba

type Config map[string]interface{}

type Sender interface {
	Send(mail *Mail, config *Config) error
}
