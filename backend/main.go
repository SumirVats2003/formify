package main

import (
	"fmt"

	"github.com/SumirVats2003/formify/backend/internal/app"
)

func main() {
	app, err := app.InitApp()
	if err != nil {
		panic(err)
	}
	fmt.Println(app)
}
