package consumer

type Consumer struct{}

func New() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Run() {

}

func (c *Consumer) Stop() {

}
