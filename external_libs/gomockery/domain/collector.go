package domain

//go:generate mockery --name=A
type A interface {
	A()
}

//go:generate mockery --name=B
type B interface {
	B()
}

//go:generate mockery --name=C
type C interface {
	CreateB(a A) B
}
