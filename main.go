package main

import (
	_ "ceph-api/docs"
	"ceph-api/system/controllers"
	"ceph-api/system/services"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
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

// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := iris.New()
	// You got full debug messages, useful when using MVC and you want to make
	// sure that your code is aligned with the Iris' MVC Architecture.
	app.Logger().SetLevel("debug")

	//jwt中间件
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte("My Secret"), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		//ErrorHandler: func(context.Context, string)

		//debug 模式
		//Debug: bool
	})
	//app.Use(jwtHandler.Serve)

	// "/pool" based mvc application.
	pool := mvc.New(app.Party("/api/pool"))
	rbd := mvc.New(app.Party("/api/rbd"))
	snap := mvc.New(app.Party("/api/snap"))
	admin := mvc.New(app.Party("/api/admin"))
	pool.Register(services.NewPoolService())
	pool.Handle(new(controllers.PoolController))
	rbd.Register(services.NewRbdService())
	rbd.Handle(new(controllers.RbdController))
	snap.Register(services.NewSnaphostService())
	snap.Handle(new(controllers.SnaphostController))
	admin.Handle(new(controllers.AdminController))

	app.Get("/api/admin/login", jwtHandler.Serve, myAuthenticatedHandler)
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
