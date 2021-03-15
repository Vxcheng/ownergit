package builtin

type fish struct {
}

func (f *fish) Swing() {}
func (f *fish) swing() error {
	panic("")
}
