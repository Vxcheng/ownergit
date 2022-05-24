package builtin

type fish struct {
}

func (f *fish) Swing() {}
func (f *fish) swing() error {
	panic("")
}

type Duck interface {
	Walk()
}

type WildDuck interface {
	Walk()
}

func Invoke(d Duck) {}

type wildDuck struct{}

func (w *wildDuck) Walk() {}

func GetWildDuck() WildDuck {
	return &wildDuck{}
}

func Do() {
	Invoke(GetWildDuck())
}
