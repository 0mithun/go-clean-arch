package main

import (
	"context"

	"github.com/0mithun/go-clean-arch/internal/app"
)

func main() {
	ctx := context.Background()
	application, err := app.NewApplication(ctx)
	if err != nil {
		panic(err)
	}

	err = application.Run(ctx)
	if err != nil {
		panic(err)
	}
}
