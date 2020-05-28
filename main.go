package main

import (
	_ "github.com/ceph/go-ceph/docs"
	"github.com/ceph/go-ceph/system/pool/controllers"
	poolService "github.com/ceph/go-ceph/system/pool/services"
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

	// ---- Serve our controllers. ----
	// Prepare our repositories and services.
	//db, err := datasource.LoadUsers(datasource.Memory)
	//if err != nil {
	//	app.Logger().Fatalf("error while loading the users: %v", err)
	//	return
	//}

	// "/pool" based mvc application.
	pool := mvc.New(app.Party("/api/pool"))
	// Add the basic authentication(admin:password) middleware
	// for the /users based requests.
	//pool.Router.Use(middleware.BasicAuth)
	// 将PoolService绑定到Controller的服务（接口）
	pool.Register(poolService.NewPoolService())
	pool.Handle(new(controllers.PoolController))

	url := swagger.URL("http://localhost:8080/swagger/doc.json") //The url pointing to API definition
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
