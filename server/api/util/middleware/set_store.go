package middleware

import (
	"server/api/sqlc"
	"server/api/util/lib"
	"server/api/util/token"

	"github.com/labstack/echo/v4"
)

func SetStore(store sqlc.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//TODO:storeを変数で管理する
			c.Set("store", store)
			return next(c)
		}
	}
}
func SetConfig(config lib.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//TODO:storeを変数で管理する
			c.Set("config", config)
			return next(c)
		}
	}
}
func SetTokenMaker(tokenMaker token.Maker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//TODO:storeを変数で管理する
			c.Set("tk",tokenMaker)
			return next(c)
		}
	}
}
