package quote

type Duck interface {
	Swing()
	swing() error
}

type fish struct {
}

func (f *fish) Swing() {}
func (f *fish) swing() error {
	panic("")
}
