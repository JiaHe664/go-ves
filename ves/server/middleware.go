package server

import (
	"github.com/HyperService-Consortium/go-ves/lib/backend/gin-helper"
	"github.com/HyperService-Consortium/go-ves/lib/backend/jwt"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/config"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/gin-contrib/cors"
	//"github.com/Myriad-Dreamin/gin-middleware/auth/privileger"
	"strconv"
)

func (srv *Server) PrepareMiddleware() bool {
	srv.jwtMW = jwt.NewMiddleWare(func() *jwt.CustomClaims {
		var cc = new(jwt.CustomClaims)
		cc.CustomField = &types.CustomFields{}
		return cc
	}, func(c controller.MContext, cc *jwt.CustomClaims) error {
		c.Set("uid", strconv.FormatInt(cc.CustomField.(*types.CustomFields).UID, 10))
		return nil
	})
	srv.jwtMW.ExpireSecond = 3600 * 24 * 7
	srv.jwtMW.RefreshSecond = 3600 * 24 * 7

	srv.routerAuthMW = controller.NewMiddleware(srv.ModelProvider.Enforcer(),
		"user:", "uid", ginhelper.MissID, ginhelper.AuthFailed)

	srv.corsMW = cors.New(cors.Config{
		//AllowAllOrigins: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowOrigins:     []string{"http://127.0.0.1:80", "https://127.0.0.1:80"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"X-Total-Count"},
		AllowCredentials: true,
	})

	srv.Module.Provide(config.ModulePath.Minimum.Middleware.JWT, srv.jwtMW)
	srv.Module.Provide(config.ModulePath.Minimum.Middleware.RouteAuth, srv.routerAuthMW)
	srv.Module.Provide(config.ModulePath.Minimum.Middleware.CORS, srv.corsMW)
	return true
}
