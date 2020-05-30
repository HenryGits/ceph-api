package main

import (
	_ "github.com/ceph/go-ceph/docs"
	"github.com/ceph/go-ceph/system/controllers"
	"github.com/ceph/go-ceph/system/services"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// @title Ceph Rest Api
// @version 1.0
// @description 基于go-ceph封装的Rest Api
// @host localhost:8080
// @BasePath /api
func main() {
	app := iris.New()
	// You got full debug messages, useful when using MVC and you want to make
	// sure that your code is aligned with the Iris' MVC Architecture.
	app.Logger().SetLevel("info")

	// "/pool" based mvc application.
	pool := mvc.New(app.Party("/api/pool"))
	rbd := mvc.New(app.Party("/api/rbd"))
	snap := mvc.New(app.Party("/api/snap"))
	pool.Register(services.NewPoolService())
	pool.Handle(new(controllers.PoolController))
	rbd.Register(services.NewRbdService())
	rbd.Handle(new(controllers.RbdController))
	snap.Register(services.NewSnaphostService())
	snap.Handle(new(controllers.SnaphostController))

	//The url pointing to API definition
	url := swagger.URL("http://localhost:8080/swagger/doc.json")
	app.Get("/swagger/{any:path}", swagger.WrapHandler(swaggerFiles.Handler, url))
	app.Run(
		// Starts the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}
