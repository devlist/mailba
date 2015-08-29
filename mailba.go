package mailba

//go:generate mockery -name Sender -output mock_mailba

type Config map[string]interface{}

type Sender interface {
	Send(mail *Mail) error
}
