package point

type logger interface {
	Debug(string)
}

type point interface {
	Execute()
	Resume()
}
