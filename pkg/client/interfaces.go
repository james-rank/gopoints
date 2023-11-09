package client

type logger interface {
	Debug(string)
}

type requester interface {
	SendRequest(buf []byte, uid string) (msg []byte, err error)
}

type subscriber interface {
	Start(string) error
	Stop() error
	Subscribe() <-chan []byte
}
