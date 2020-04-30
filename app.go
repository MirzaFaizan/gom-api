package main

import (
	"os"

	R "github.com/mirzafaizan/gom-api/controllers"

	"github.com/kataras/iris"

	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET Default Endpoint
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.JSON(context.Map{"message": "Welcome to my API"})
	})

	//API end points
	api := app.Party("/api")
	{
		api.Post("/signup", R.CreateUser)
		api.Post("/login", R.GetUser)
		api.Get("/getusers", R.GetAllUsers)
		api.Get("/users/{msisdn: string}", R.GetUser)
	}

	app.Run(iris.Addr(os.Getenv("PORT")), iris.WithoutServerError(iris.ErrServerClosed))
}
