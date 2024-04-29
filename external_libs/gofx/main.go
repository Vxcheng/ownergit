package main

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

func main() {
	demo2()
}

func demo2() {
	type a struct {
		name string
	}
	type in struct {
		fx.In

		A *a `name:"foo"`
	}
	newA := func() *a {
		return &a{name: "test"}
	}

	var xx in
	app := fx.New(
		fx.Provide(
			fx.Annotated{
				Name:   "foo",
				Target: newA,
			},
		),
		fx.Populate(&xx),
	)

	fmt.Println(xx)
	defer app.Done()

}

func demo1() {
	// Some external module that provides a user name.
	type Username string
	UserModule := fx.Provide(func() Username { return "john" })

	// We want to use Fx to wire up our constructors, but don't actually want to
	// run the application - we just want to yank out the user name.
	//
	// This is common in unit tests, and is even easier with the fxtest
	// package's RequireStart and RequireStop helpers.
	var user Username
	app := fx.New(
		UserModule,
		fx.NopLogger, // silence test output
		fx.Populate(&user),
	)
	if err := app.Start(context.Background()); err != nil {
		panic(err)
	}
	defer app.Stop(context.Background())

	fmt.Println(user)

}
