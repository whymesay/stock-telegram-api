package command

type Command interface {
	exec(args ...interface{})
}
