package types

type ILogger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
	Warn(msg string)
	Fatal(msg string)
}

type ISuggarEngine interface {
	Warn(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
	Fatal(args ...interface{})
}
