package producer

type Producer struct{}

func New() *Producer {
	// init
	return &Producer{}
}

func (s *Producer) Send(topic string, msg interface{}) error {
	return nil
}
