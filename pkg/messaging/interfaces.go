// Package messaging is a package that defines the wrappers around mangos.
package messaging

type logger interface {
	Debug(string)
}

type socket interface {
	Close() error
	Dial(string) error
	Listen(string) error
	Recv() ([]byte, error)
	Send([]byte) error
	SetOption(string, interface{}) error
}
