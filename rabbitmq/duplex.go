package broker

type Duplex struct {
	*Consumer
	*Producer
}
