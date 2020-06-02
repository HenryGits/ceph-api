package main

import (
	_ "github.com/ceph/go-ceph/docs"
	"github.com/ceph/go-ceph/system/controllers"
	"github.com/ceph/go-ceph/system/services"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// @title Ceph Rest Api
// @version 1.0
// @description 基于go-ceph封装的Rest Api
// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := iris.New()
	// You got full debug messages, useful when using MVC and you want to make
	// sure that your code is aligned with the Iris' MVC Architecture.
	app.Logger().SetLevel("debug")

	//jwt中间件
	j := jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return "My Secret", nil
		},

		// Extract by the "token" url.
		// There are plenty of options.
		// The default jwt's behavior to extract a token value is by
		// the `Authorization: Bearer $TOKEN` header.
		Extractor: jwt.FromParameter("token"),
		// When set, the middleware verifies that tokens are
		// signed with the specific signing algorithm
		// If the signing method is not constant the `jwt.Config.ValidationKeyGetter` callback
		// can be used to implement additional checks
		// Important to avoid security issues described here:
		// https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})

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

	app.Get("/api/admin/auth", j.Serve, myAuthenticatedHandler)
	//The url pointing to API definition
	config := &swagger.Config{
		URL: "http://localhost:8080/swagger/doc.json",
	}
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
	app.Run(
		// Starts the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}

func myAuthenticatedHandler(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)
	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")
	foobar := user.Claims.(jwt.MapClaims)
	for key, value := range foobar {
		ctx.Writef("%s = %s", key, value)
	}
}
